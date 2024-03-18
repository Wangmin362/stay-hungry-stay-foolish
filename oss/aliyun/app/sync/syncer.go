package sync

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/fsnotify/fsnotify"
	"os"
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
		if err := createBucket(bucketName, client); err != nil {
			return nil, nil
		}
	}

	if err = setRefer(client, bucketName); err != nil {
		return nil, fmt.Errorf("set refer error: %w", err)
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

	s := &Syncer{
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
	if err := s.cacheAllAliyunObjs(); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	// 全量同步一次
	if err := s.SyncDirPic(syncDir); err != nil {
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

func (s *Syncer) cacheAllAliyunObjs() error {
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
