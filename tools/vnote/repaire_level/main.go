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

const (
	levelPattern string = `(#+) (\d[\d|\.]*) (\d[\d|\.]*? )`

	picPattern       string = `\!\[.*?\]\((.*?)(?: \".*?)?(?: =.*?)?\)`
	linkPattern      string = `\[.*?\]\((.*?)(?: \".*?)?(?: =.*?)?\)`
	codeBlockPattern string = "```(.*+?)```"
	inlinePattern    string = "`(.*+?)`"
)

// 修改字体颜色的格式
func main() {
	path := "D:/Notebook/Vnote"
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

		// TODO 优化以下，只需要读取一次文件、写入一次文件
		if err = RepairMarkdown(path); err != nil {
			return err
		}

		return RepairHighLight(path)
	})
}

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
	re := regexp.MustCompile(levelPattern)
	match := re.FindAllSubmatch(rawData, -1)
	for _, group := range match {
		raw := string(group[0])
		level := string(group[1])
		num := string(group[2])

		target := fmt.Sprintf(`%s %s `, level, num)
		fileData = strings.ReplaceAll(fileData, raw, target)
	}

	if bytes.Equal(rawData, []byte(fileData)) {
		return nil
	}

	if err = os.WriteFile(path, []byte(fileData), os.ModePerm); err != nil {
		return err
	}

	fmt.Printf("%s文件标题级别处理完成\n", path)

	return nil
}

const (
	lineHighLightPattern string = `(#+) (\d[\d|\.]*) (.*?\n)`
)

func RepairHighLight(path string) error {
	// 当前同步的文件必须是以.md结尾，也就是当前文件必须是一个markdown格式的文件才进行修改
	if filepath.Ext(path) != ".md" {
		return nil
	}

	rawData, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	fileData := string(bytes.Clone(rawData))
	re := regexp.MustCompile(lineHighLightPattern)
	match := re.FindAllSubmatch(rawData, -1)
	for _, group := range match {
		raw := string(group[0])
		level := string(group[1])
		num := string(group[2])
		title := string(group[3])

		title = strings.ReplaceAll(title, "`", "")
		title = strings.ReplaceAll(title, "*", "")

		target := fmt.Sprintf(`%s %s %s`, level, num, title)
		fileData = strings.ReplaceAll(fileData, raw, target)
	}

	if bytes.Equal(rawData, []byte(fileData)) {
		return nil
	}

	if err = os.WriteFile(path, []byte(fileData), os.ModePerm); err != nil {
		return err
	}

	fmt.Printf("%s文件标题星号、反引号处理完成\n", path)

	return nil
}

func ConvertEnglishToInline(data []byte, path string) ([]byte, error) {
	// 1、排除博客目录以外的文件
	// 2、创建一个map用于保存图片、链接、代码块、行内语法
	// 3、替换图片语法
	// 4、替换链接语法
	// 5、替换代码块语法
	// 6、替换行内语法
	// 7、把所有的英文替换为行内语法格式
	// 8、恢复图片、链接、代码块、行内语法格式
	// 9、返回
	return nil, nil
}
