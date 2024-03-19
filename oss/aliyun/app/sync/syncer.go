package sync

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"

	"github.com/golang/demo/oss/aliyun/app/sync/cache"
)

func NewSyncer(syncDir, endpoint, bucketName, ossId, ossSecret string) (*syncer, error) {
	// 创建阿里云OSS客户端
	client, err := oss.New(fmt.Sprintf("https://%s", endpoint), ossId, ossSecret)
	if err != nil {
		return nil, fmt.Errorf("create aliyun oss client error:%w", err)
	}

	// 判断指定的桶是否存在
	exist, err := client.IsBucketExist(bucketName)
	if err != nil {
		return nil, fmt.Errorf("query %s bucket exist error:%w", bucketName, err)
	}

	// 如果桶不存在，就创建这个桶
	if !exist {
		if err := CreateStandardLRSReadPublicBucket(bucketName, client); err != nil {
			return nil, fmt.Errorf("create %s bucket error: %w", bucketName, err)
		}
	}

	// 为当前桶设置防盗链，防止流量盗刷
	if err = SetReferer(client, bucketName,
		[]string{"*.jianshu.com"},
		[]string{"*.baidu.com"},
	); err != nil {
		return nil, fmt.Errorf("set bucket referer error: %w", err)
	}

	// 获取当前桶
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, fmt.Errorf("get %s bucket error:%w", bucketName, err)
	}

	// 创建一个文件监听器，当文件、目录发生变化时，我们可以及时知道，而不必每次循环迭代扫描所有文件
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, fmt.Errorf("create file watcher:%w", err)
	}

	// 监听指定的同步目录的文件变化
	if err = watcher.Add(syncDir); err != nil {
		return nil, fmt.Errorf("watch dir %s error:%w", syncDir, err)
	}

	s := &syncer{
		client:     client,
		bucket:     bucket,
		endpoint:   endpoint,
		bucketName: bucketName,
		ossId:      ossId,
		ossSecret:  ossSecret,
		syncDir:    syncDir,
		imageDir:   syncImageDir, // 设置仅针对某些特殊的目录上传文件
		Cache:      cache.NewCache(),
		fsWatcher:  watcher,
	}

	// 先把当前bucket中所有缓存的文件名存储起来
	if err := s.cacheAllAliOSSObjs(); err != nil {
		return nil, err
	}

	// 全量同步一次
	if err := s.syncDirPic(syncDir); err != nil {
		return nil, err
	}

	return s, nil
}

type Empty struct{}

type syncer struct {
	client *oss.Client
	bucket *oss.Bucket

	endpoint   string
	bucketName string
	ossId      string
	ossSecret  string
	syncDir    string // 需要同步的目录
	imageDir   string // 如果设置，那么仅同步名字为指定目录下的文件，否则同步所有文件

	fsWatcher *fsnotify.Watcher

	markdownLock sync.Mutex

	*cache.Cache
}

func (s *syncer) cacheAllAliOSSObjs() error {
	continueToken := ""
	for {
		lsRes, err := s.bucket.ListObjectsV2(oss.ContinuationToken(continueToken))
		if err != nil {
			return err
		}

		// 打印列举结果。默认情况下，一次返回100条记录。
		for _, obj := range lsRes.Objects {
			s.CacheObj(obj.Key)
		}

		if lsRes.IsTruncated {
			continueToken = lsRes.NextContinuationToken
		} else {
			break
		}
	}

	return nil
}

func (s *syncer) syncDirPic(syncDir string) error {
	return filepath.Walk(syncDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if syncDir == path {
			return nil
		}

		if info.IsDir() {
			if err = s.fsWatcher.Add(path); err != nil { // 子目录也需要监视
				return fmt.Errorf("add watch %s dir error: %w", path, err)
			}
			return nil
		}

		// 当前文件是普通文件，直接上传到阿里云
		if err = s.saveToAliOss(path); err != nil {
			return err
		}
		return nil
	})
}

func (s *syncer) saveToAliOss(path string) error {
	// 当前文件路径必须包含指定的路径才是需要同步的文件,否则直接跳过
	if !strings.Contains(path, s.imageDir) {
		return nil
	}

	// 如果目录不正确，直接跳过
	index := strings.Index(path, s.syncDir)
	if index < 0 {
		return errors.Errorf("%s目录不正确，基础目录不是%s", path, s.syncDir)
	}

	// 获取子路径
	realPath := path[len(s.syncDir)+1:]
	dstBucketKey := ConvertWindowDirToLinuxDir(realPath)

	if s.ObjExist(dstBucketKey) { // 说明当前文件已经同步
		return nil
	}

	if err := SaveToAliOSS(path, dstBucketKey, s.bucket); err != nil {
		if errors.Is(err, FileZeroSize) {
			return nil
		}
		return fmt.Errorf("同步%s文件到阿里云错误: %w", path, err)
	}
	s.CacheObj(dstBucketKey)

	fmt.Printf("同步%s文件到阿里云%s成功!\n", path, dstBucketKey)
	return nil
}

func (s *syncer) moveFile(dst, src string) error {
	if err := MoveFile(dst, src, s.bucket); err != nil {
		return err
	}

	s.Replace(dst, src)
	return nil
}

func (s *syncer) Run() {
	defer s.fsWatcher.Close()

	// 增量同步
	go s.watchDirTask()

	for {
		if err := s.syncDirPic(s.syncDir); err != nil {
			fmt.Printf("%s\n", err)
		}

		if err := s.replaceDirPic(s.syncDir); err != nil {
			fmt.Printf("%s\n", err)
		}
		time.Sleep(10 * time.Minute)
	}
}
