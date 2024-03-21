package sync

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	pattern string = `\!\[.*?\]\((.*?)(?: \".*?)?(?: =.*?)?\)`
)

// 在markdown中一共有如下四种引用格式
// ![](addr)
// ![](addr "sdfsdf")
// ![](addr "sdfsdf" =1220x)

func (s *syncer) replaceMarkdownPicRef(mdPath string) error {
	// 当前同步的文件必须是以.md结尾，也就是当前文件必须是一个markdown格式的文件才进行修改
	if filepath.Ext(mdPath) != ".md" {
		return nil
	}

	var err error
	if mdPath, err = filepath.Abs(mdPath); err != nil {
		return err
	}

	s.markdownLock.Lock()
	rawData, err := os.ReadFile(mdPath)
	if err != nil {
		s.markdownLock.Unlock()
		return err
	}
	s.markdownLock.Unlock()

	dir := filepath.Dir(mdPath)
	if strings.Index(dir, s.syncDir) < 0 {
		return fmt.Errorf("路径不正确:%s", mdPath)
	}

	if dir != s.syncDir {
		dir = dir[len(s.syncDir)+1:]
		dir = ConvertWindowDirToLinuxDir(dir)
	} else {
		dir = ""
	}

	// 文件数据
	fileData := string(bytes.Clone(rawData))
	re := regexp.MustCompile(pattern)
	match := re.FindAllSubmatch(rawData, -1)
	for _, group := range match {
		markdownPic := string(group[0])
		imagePath := string(group[1])
		aliOSSKey := fmt.Sprintf("%s/%s", dir, imagePath)
		if dir == "" {
			aliOSSKey = fmt.Sprintf("%s", imagePath)
		}

		// 如果当前图片的引用路径已经就是阿里云的路径，说明不需要替换；否则说明是本地路径，需要进行替换
		aliOssUrl := fmt.Sprintf("https://%s.%s", s.bucketName, s.endpoint)
		if !strings.Contains(imagePath, aliOssUrl) {
			if !s.ObjExist(aliOSSKey) { // 图片还没有上传，那就先尝试上传一次
				picPath := filepath.Join(filepath.Dir(mdPath), imagePath)
				_ = s.saveToAliOss(picPath) // 不关心上传失败没有
			}

			if s.ObjExist(aliOSSKey) { // 如果这次查询，已经上传了图片，那就直接替换
				repAddr := fmt.Sprintf("![](%s/%s)", aliOssUrl, aliOSSKey)
				fileData = strings.ReplaceAll(fileData, markdownPic, repAddr)
			}
		} else {
			currOssKey := imagePath[len(aliOssUrl)+1:]
			rightOssKey := fmt.Sprintf("%s/%s/%s", dir, s.imageDir, filepath.Base(imagePath))
			if dir == "" {
				rightOssKey = fmt.Sprintf("%s/%s", s.imageDir, filepath.Base(imagePath))
			}
			rightOssUrl := fmt.Sprintf("![](%s/%s)", aliOssUrl, rightOssKey)

			var repAddr string
			if currOssKey == rightOssKey {
				repAddr = markdownPic
			} else { // 如果不相等，说明文件移动位置了
				if s.ObjExist(rightOssKey) { // 如果这次查询，图片已经存在
					repAddr = rightOssUrl
				} else { // 说明当前没有上传，那就先拷贝一份
					if err = s.moveFile(rightOssKey, currOssKey); err != nil {
						// 如果出错了，就不修正位置
						repAddr = markdownPic
					} else { // 否则，就修正引用位置
						if s.ObjExist(rightOssKey) {
							repAddr = rightOssUrl
						} else {
							repAddr = markdownPic
						}
					}
				}
			}
			if markdownPic != repAddr {
				fileData = strings.ReplaceAll(fileData, markdownPic, repAddr)
			}
		}
	}

	// 说明文件没有改动，直接退出
	if bytes.Equal(rawData, []byte(fileData)) {
		return nil
	}

	// 替换文件内容
	s.markdownLock.Lock()
	defer s.markdownLock.Unlock()
	if err = os.WriteFile(mdPath, []byte(fileData), os.ModePerm); err != nil {
		return err
	}

	return nil
}

func (s *syncer) replaceDirPic(syncDir string) error {
	return filepath.Walk(syncDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if syncDir == path {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		if err = s.replaceMarkdownPicRef(path); err != nil {
			return fmt.Errorf("替换%s文件图片路径错误: %w", path, err)
		}
		return nil
	})
}
