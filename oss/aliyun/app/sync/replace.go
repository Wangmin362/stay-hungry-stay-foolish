package sync

import (
	"bytes"
	"fmt"
	"io"
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

func (s *Syncer) replaceMarkdownPicRef(mdPath string) error {
	if filepath.Ext(mdPath) != ".md" {
		return nil
	}

	var err error
	if mdPath, err = filepath.Abs(mdPath); err != nil {
		return err
	}

	file, err := os.OpenFile(mdPath, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()
	rawData, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	dir := filepath.Dir(mdPath)
	index := strings.Index(dir, s.syncDir)
	if index < 0 {
		return fmt.Errorf("路径不正确:%s", mdPath)
	}
	p := dir[len(s.syncDir)+1:]
	split := strings.Split(p, "\\")
	p = strings.Join(split, "/")

	fileData := string(rawData)
	re := regexp.MustCompile(pattern)
	match := re.FindAllSubmatch(rawData, -1)
	for _, group := range match {
		if !strings.Contains(string(group[1]), fmt.Sprintf("https://%s.%s", s.bucketName, s.endpoint)) {
			aliossPic := fmt.Sprintf("%s/%s", p, string(group[1]))
			if _, ok := s.cacheObjs[aliossPic]; !ok { // 图片还没有上传，那就先尝试上传一次
				picPath := fmt.Sprintf("%s/%s", dir, string(group[1]))
				_ = s.syncFileToAliyun(picPath) // 不关心上传失败没有
			}

			if _, ok := s.cacheObjs[aliossPic]; ok { // 如果这次查询，已经上传了图片，那就直接替换
				repAddr := fmt.Sprintf("![](https://%s.%s/%s/%s)", s.bucketName, s.endpoint, p, group[1])
				fileData = strings.ReplaceAll(fileData, string(group[0]), repAddr)
			}
		} else {
			repAddr := fmt.Sprintf("![](%s)", group[1])
			fileData = strings.ReplaceAll(fileData, string(group[0]), repAddr)
		}
	}
	if bytes.Equal(rawData, []byte(fileData)) {
		return nil
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}
	_, err = file.Write([]byte(fileData))
	if err != nil {
		return err
	}

	return nil
}

func (s *Syncer) ReplaceDirPic(syncDir string) error {
	return filepath.Walk(syncDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if syncDir == path {
			return nil
		}

		if info.IsDir() {
			return s.ReplaceDirPic(path)
		}

		if err = s.replaceMarkdownPicRef(path); err != nil {
			return fmt.Errorf("替换%s文件图片路径错误: %w", path, err)
		}
		return nil
	})
}
