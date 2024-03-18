package sync

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"os"
	gpath "path"
	"path/filepath"
	"strings"
	"time"
)

func NewSyncer(syncDir, endpoint, bucketName, ossId, ossSecret string) (*syncer, error) {

	client, err := oss.New(fmt.Sprintf("https://%s", endpoint), ossId, ossSecret)
	if err != nil {
		return nil, fmt.Errorf("create aliyun oss client error:%w", err)
	}

	exist, err := client.IsBucketExist(bucketName)
	if err != nil {
		return nil, fmt.Errorf("query %s bucket exist error:%w", bucketName, err)
	}

	if !exist {
		if err := CreateStandardLRSReadPublicBucket(bucketName, client); err != nil {
			return nil, nil
		}
	}

	if err = SetReferer(client, bucketName,
		[]string{"*.jianshu.com"},
		[]string{"*.baidu.com"},
	); err != nil {
		return nil, fmt.Errorf("set bucket referer error: %w", err)
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, fmt.Errorf("get %s bucket error:%w", bucketName, err)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, fmt.Errorf("create file watcher:%w", err)
	}

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
		imageDir:   syncImageDir,
		cacheObjs:  make(map[string]Empty),
		fsWatcher:  watcher,
	}

	// 先把当前bucket中所有缓存的文件名存储起来
	if err := s.cacheAllAliOSSObjs(); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	// 全量同步一次
	if err := s.syncDirPic(syncDir); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
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

	cacheObjs map[string]Empty
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
			s.cacheObjs[obj.Key] = Empty{}
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
			// 当前目录是想要同步的目录才考虑上传阿里云
			if info.Name() == s.imageDir {
				return s.syncDirPic(path)
			}

			if err = s.fsWatcher.Add(path); err != nil { // 子目录也需要监视
				return fmt.Errorf("watch %s dir error: %w", path, err)
			}
		}

		if !strings.Contains(path, s.imageDir) { // 必须是目标目录才同步
			return nil
		}

		index := strings.Index(path, s.syncDir)
		if index < 0 {
			return errors.Errorf("%s目录不正确，基础目录不是%s", path, s.syncDir)
		}

		realPath := path[len(s.syncDir)+1:]
		split := strings.Split(realPath, "\\")
		dstBucketKey := gpath.Join(split...)

		_, ok := s.cacheObjs[dstBucketKey]
		if ok {
			fmt.Printf("@@@已经同步%s文件到阿里云,访问路径为:%s\n", path, fmt.Sprintf(url, s.bucketName, s.endpoint, dstBucketKey))
			return nil
		}

		if err = SaveToAliOSS(path, dstBucketKey, s.bucket); err != nil {
			return fmt.Errorf("同步%s文件到阿里云错误: %w", path, err)
		}

		fmt.Printf("同步%s文件到阿里云%s成功!\n", path, dstBucketKey)
		s.cacheObjs[dstBucketKey] = Empty{}
		return nil
	})
}

func (s *syncer) Run() {
	defer s.fsWatcher.Close()

	// 增量同步
	go s.watchDir()

	for {
		if err := s.ReplaceDirPic(s.syncDir); err != nil {
			fmt.Printf("%s\n", err)
		}
		time.Sleep(10 * time.Minute)
	}
}
