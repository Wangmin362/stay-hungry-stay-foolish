package sync

import (
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
)

func (s *syncer) watchDirTask() {
	// 目录（增、删除、重命名）、文件（增、删除、重命名、修改）
	for {
		select {
		case event, ok := <-s.fsWatcher.Events:
			if !ok {
				return
			}

			currPath := event.Name
			info, err := os.Stat(currPath)
			if err != nil { // 文件很有可能被删除了，直接忽略即可；除此之外，一般不会报错
				continue
			}
			if info.IsDir() { // 目录发生了变更
				if event.Has(fsnotify.Create) || event.Has(fsnotify.Rename) {
					if err = s.fsWatcher.Add(currPath); err != nil {
						log.Printf("watch %s dir error: %s\n", currPath, err)
					}
				}
				if event.Has(fsnotify.Remove) {
					if err = s.fsWatcher.Remove(currPath); err != nil {
						log.Printf("remote %s dir from fswatcher error: %s\n", currPath, err)
					}
				}
				continue
			}

			// 说明当前路径是文件

			if event.Has(fsnotify.Remove) || event.Has(fsnotify.Chmod) {
				// 文件修改权限、删除，直接忽略
				continue
			}

			if err = s.saveToAliOss(currPath); err != nil {
				log.Printf("%s\n", err)
				continue
			}

			// 替换markdown文件的内容
			if err = s.replaceMarkdownPicRef(currPath); err != nil {
				log.Printf("%s\n", err)
				continue
			}
		case err, ok := <-s.fsWatcher.Errors:
			if !ok {
				return
			}
			log.Printf("error:%s\n", err)
		}
	}
}
