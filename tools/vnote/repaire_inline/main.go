package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// 修改字体颜色的格式
func main() {
	path := "D:/Notebook/Vnote/Blog"
	if err := RepairDir(path); err != nil {
		log.Fatal(err)
	}
}

func RepairDir(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if dir == path {
			return nil
		}

		if info.IsDir() {
			return nil // 目的直接跳过
		}

		return RepairMarkdown(path)
	})
}

const (
	pattern string = "`([0-9][0-9\\.]*?\\.)`"
)

func RepairMarkdown(path string) error {
	// 当前同步的文件必须是以.md结尾，也就是当前文件必须是一个markdown格式的文件才进行修改
	if filepath.Ext(path) != ".md" {
		return nil
	}

	rawData, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	fileData := string(bytes.Clone(rawData))
	re := regexp.MustCompile(pattern)
	match := re.FindAllSubmatch(rawData, -1)
	for _, group := range match {
		raw := string(group[0])
		target := string(group[1])

		fileData = strings.ReplaceAll(fileData, raw, target)
	}

	if bytes.Equal(rawData, []byte(fileData)) {
		return nil
	}

	if err = os.WriteFile(path, []byte(fileData), os.ModePerm); err != nil {
		return err
	}

	fmt.Printf("%s文件处理完成\n", path)

	return nil
}
