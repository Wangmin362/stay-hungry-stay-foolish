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

func RepairMarkdown(path string) error {
	// 当前同步的文件必须是以.md结尾，也就是当前文件必须是一个markdown格式的文件才进行修改
	if filepath.Ext(path) != ".md" {
		return nil
	}

	// 只处理博客目录
	if !strings.Contains(path, "D:/Notebook/Vnote/Blog") && !strings.Contains(path, "D:\\Notebook\\Vnote\\Blog") {
		return nil
	}

	rawData, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	fileData := string(bytes.Clone(rawData))

	for _, ptn := range []string{"<`?font.+?>", "<`?tr.+?>", "<`?td.+?>", "</`?font`?>", "</`?td`?>", "</`?tr`?>"} {
		re := regexp.MustCompile(ptn)
		match := re.FindAllSubmatch(rawData, -1)
		for _, group := range match {
			raw := string(group[0])

			fileData = strings.ReplaceAll(fileData, raw, "")
		}
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
