package sync

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	gpath "path"
	"strings"
)

func (s *syncer) watchDir() {
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
				if !strings.Contains(currPath, s.imageDir) { // 必须是目标目录才同步
					continue
				}

				index := strings.Index(currPath, s.syncDir)
				if index < 0 {
					fmt.Printf("%s目录不正确，基础目录不是%s", currPath, s.syncDir)
					continue
				}

				realPath := currPath[len(s.syncDir)+1:]
				split := strings.Split(realPath, "\\")
				dstBucketKey := gpath.Join(split...)

				_, ok := s.cacheObjs[dstBucketKey]
				if ok {
					fmt.Printf("@@@已经同步%s文件到阿里云,访问路径为:%s\n", currPath, fmt.Sprintf(url, s.bucketName, s.endpoint, dstBucketKey))
					continue
				}

				if err = SaveToAliOSS(currPath, dstBucketKey, s.bucket); err != nil {
					fmt.Printf("%s\n", err)
					continue
				}

				fmt.Printf("同步%s文件到阿里云%s成功!\n", currPath, dstBucketKey)
				s.cacheObjs[dstBucketKey] = Empty{}
			}

			if event.Has(fsnotify.Create) || event.Has(fsnotify.Rename) || event.Has(fsnotify.Write) {
				if err = s.replaceMarkdownPicRef(currPath); err != nil {
					fmt.Printf("%s\n", err)
				}
			}
		case err, ok := <-s.fsWatcher.Errors:
			if !ok {
				return
			}
			log.Printf("error:%s\n", err)
		}
	}
}
