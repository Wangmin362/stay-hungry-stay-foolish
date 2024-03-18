package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/golang/demo/oss/aliyun/app/sync"
	"log"
	"os"
	gpath "path"
	"path/filepath"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg/errors"
)

const (
	EndpointKey  = "EndpointKey"
	BucketKey    = "BucketKey"
	OssIDKey     = "OSS_ACCESS_KEY_ID"
	OssSecretKey = "OSS_ACCESS_KEY_SECRET"
	SyncDirKey   = "SyncDirKey"
)

func getEnvVar(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", errors.Errorf("environment variable %s key not found", key)
	}

	return value, nil
}


// 1、本地文件删除之后暂时不考虑删除云端的文件，保留备份，以免后面还需要
// TODO 2、考虑目录的重命名暂时不处理，后续写一个定时任务，直接清楚阿里云OSS中没有使用的图片
// TODO 3、修正文件移动后，路径不对的问题
// TODO 如何保证图片的安全？ 防止其他人胡乱使用？
func main() {
	var err error
	syncDir, err := getEnvVar(SyncDirKey)
	if err != nil {
		fmt.Printf("%s\n", syncDir)
		os.Exit(1)
	}

	endpoint, err := getEnvVar(EndpointKey)
	if err != nil {
		fmt.Printf("%s\n", syncDir)
		os.Exit(1)err
	}
	bucketName, err := getEnvVar(BucketKey)
	if err != nil {
		fmt.Printf("%s\n", syncDir)
		os.Exit(1)
	}
	ossId, err := getEnvVar(OssIDKey)
	if err != nil {
		fmt.Printf("%s\n", syncDir)
		os.Exit(1)
	}
	ossSecret, err := getEnvVar(OssSecretKey)
	if err != nil {
		fmt.Printf("%s\n", syncDir)
		os.Exit(1)
	}

	syncer, err := sync.NewSyncer(syncDir, endpoint, bucketName, ossId, ossSecret)
	if err != nil {
		fmt.Printf("%s\n", syncDir)
		os.Exit(1)
	}
	defer syncer.fsWatcher.Close()

	// 增量同步
	go syncer.watchDir()

	for {
		if err := syncer.ReplaceDirPic(syncDir); err != nil {
			fmt.Printf("%s\n", err)
		}
		time.Sleep(10 * time.Minute)
	}
}




func (s *Syncer) SyncDirPic(syncDir string) error {
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
				return s.SyncDirPic(path)
			}

			if err = s.fsWatcher.Add(path); err != nil { // 子目录也需要监视
				return fmt.Errorf("watch %s dir error: %w", path, err)
			}
		}

		if err = s.syncFileToAliyun(path); err != nil {
			return fmt.Errorf("同步%s文件到阿里云错误: %w", path, err)
		}
		return nil
	})
}

func (s *Syncer) syncFileToAliyun(imgPath string) error {
	if !strings.Contains(imgPath, s.imageDir) { // 必须是目标目录才同步
		return nil
	}

	index := strings.Index(imgPath, s.syncDir)
	if index < 0 {
		return errors.Errorf("%s目录不正确，基础目录不是%s", imgPath, s.syncDir)
	}

	info, err := os.Stat(imgPath)
	if err != nil {
		return fmt.Errorf("statistic %s path error:%w", imgPath, err)
	}
	if info.IsDir() { // 忽略目录
		return nil
	}

	// 如果当前图片的大小为0，暂时先不同步
	if info.Size() <= 0 {
		return nil
	}

	realPath := imgPath[len(s.syncDir)+1:]
	split := strings.Split(realPath, "\\")
	dstBucketKey := gpath.Join(split...)

	_, ok := s.cacheObjs[dstBucketKey]
	if ok {
		fmt.Printf("@@@已经同步%s文件到阿里云,访问路径为:%s\n", imgPath, fmt.Sprintf(url, s.bucketName, s.endpoint, dstBucketKey))
		return nil
	}
	if err := s.bucket.PutObjectFromFile(dstBucketKey, imgPath); err != nil {
		return fmt.Errorf("保存 %s到阿里云失败:%w", imgPath, err)
	}
	fmt.Printf("同步%s文件到阿里云%s成功!\n", imgPath, dstBucketKey)
	s.cacheObjs[dstBucketKey] = Empty{}

	return nil

}

func (s *Syncer) watchDir() {
	// 目录（增、删除、重命名）、文件（增、删除、重命名、修改）
	for {
		select {
		case event, ok := <-s.fsWatcher.Events:
			if !ok {
				return
			}

			currPath := event.Name
			info, err := os.Stat(currPath)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if info.IsDir() {
				// 如果当前路径是目录的增删改查
				if event.Has(fsnotify.Create) || event.Has(fsnotify.Rename) {
					if err = s.fsWatcher.Add(currPath); err != nil {
						fmt.Printf("watch %s dir error: %s", currPath, err)
					}
				}
				if event.Has(fsnotify.Remove) {
					if err = s.fsWatcher.Remove(currPath); err != nil {
						fmt.Printf("remote %s dir from fswatcher error: %s", currPath, err)
					}
				}
				continue
			}

			// 说明当前路径是文件

			// 一个文件同步出错了，直接忽略
			if event.Has(fsnotify.Create) || event.Has(fsnotify.Rename) || event.Has(fsnotify.Write) {
				if err = s.syncFileToAliyun(currPath); err != nil {
					fmt.Printf("%s\n", err)
					continue
				}
			}

			if err = s.replaceMarkdownPicRef(currPath); err != nil {
				fmt.Printf("%s\n", err)
			}
		case err, ok := <-s.fsWatcher.Errors:
			if !ok {
				return
			}
			log.Printf("error:%s\n", err)
		}
	}
}
