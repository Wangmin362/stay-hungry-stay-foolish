package main

import (
	"bytes"
	"fmt"
	"github.com/dlclark/regexp2"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	// levelPattern 把1.1.2修复为1.1.2.这种格式
	levelPattern string = `(#+) (\d[\d|\.]*) (\d[\d|\.]*? )`

	picPattern       string = `\!\[.*?\]\((.*?)(?: \".*?)?(?: =.*?)?\)`
	linkPattern      string = `\[.*?\]\((.*?)(?: \".*?)?(?: =.*?)?\)`
	codeBlockPattern string = "(?s)```.+?```"
	inlinePattern    string = "`.+?`"
	englishPattern   string = `[a-zA-Z0-9][a-zA-Z0-9/ ]+(?<! )`

	lineHighLightPattern string = `#{1,6} +.+`
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

		// 当前同步的文件必须是以.md结尾，也就是当前文件必须是一个markdown格式的文件才进行修改
		if filepath.Ext(path) != ".md" {
			return nil
		}

		var data []byte
		data, err = os.ReadFile(path)
		if err != nil {
			return err
		}
		rawData := bytes.Clone(data)

		// 用于修复markdown标题，把1.1.2修复为1.1.2.这种格式
		data, _ = RepairLevelFormat(data)

		// 用于把Markdown中所有的英文设置为行内语法
		data, err = ConvertEnglishToInline(data, path)
		if err != nil {
			return err
		}

		if bytes.Equal(data, rawData) {
			return nil
		}

		if err = os.WriteFile(path, data, os.ModePerm); err != nil {
			return err
		}

		fmt.Printf("%s文件处理完成\n", path)

		return nil
	})
}

// RepairLevelFormat 用于修复markdown标题，把1.1.2修复为1.1.2.这种格式
func RepairLevelFormat(data []byte) ([]byte, error) {
	fileData := string(bytes.Clone(data))
	re := regexp.MustCompile(levelPattern)
	match := re.FindAllSubmatch(data, -1)
	for _, group := range match {
		raw := string(group[0])
		level := string(group[1])
		num := string(group[2])

		target := fmt.Sprintf(`%s %s `, level, num)
		fileData = strings.ReplaceAll(fileData, raw, target)
	}

	return []byte(fileData), nil
}

// RepairLevelHighLight 用于把标题的反引号、加粗去除掉
func RepairLevelHighLight(data []byte) ([]byte, error) {
	fileData := string(bytes.Clone(data))
	re := regexp.MustCompile(lineHighLightPattern)
	match := re.FindAllSubmatch(data, -1)
	for _, group := range match {
		raw := string(group[0])

		title := strings.ReplaceAll(raw, "`", "")
		title = strings.ReplaceAll(title, "*", "")

		fileData = strings.ReplaceAll(fileData, raw, title)
	}

	return []byte(fileData), nil
}

// OrderMapping 按照序号重新排序
func OrderMapping(data []byte, mMap map[string]string) (map[string]string, []byte) {
	fileData := string(bytes.Clone(data))
	re := regexp.MustCompile("@~[0-9]+~@")
	match := re.FindAllSubmatch(data, -1)
	orderMap := make(map[string]string, len(mMap))
	for idx, group := range match {
		raw := string(group[0]) // 原本的Key

		target := fmt.Sprintf("@@%d@@", idx) // 按顺序的key
		fileData = strings.ReplaceAll(fileData, raw, target)
		orderMap[target] = mMap[raw] // 重新映射
	}

	return orderMap, []byte(fileData)
}

func ConvertEnglishToInline(data []byte, path string) ([]byte, error) {
	fileData := string(bytes.Clone(data))
	// 1、排除博客目录以外的文件
	if !strings.Contains(path, "D:/Notebook/Vnote/Blog") && !strings.Contains(path, "D:\\Notebook\\Vnote\\Blog") {
		return data, nil
	}

	// QA文件可以不管，没有必要
	if strings.Contains(path, "QA.md") {
		return data, nil
	}

	// 2、创建一个map用于保存图片、链接、代码块、行内语法，替换为@~%d~@的格式
	mMap := make(map[string]string)
	idx := 0

	// 匹配所有的图片
	for _, pattern := range []string{picPattern, linkPattern, codeBlockPattern, inlinePattern} {
		re := regexp2.MustCompile(pattern, 0)
		match, err := re.FindStringMatch(fileData)
		if err != nil {
			return nil, err
		}
		ma := func(match *regexp2.Match) {
			groups := match.Groups()
			raw := string(groups[0].Captures[0].Runes()) // 匹配到的字符串

			target := fmt.Sprintf("@~%d~@", idx)
			mMap[target] = raw
			fileData = strings.Replace(fileData, raw, target, 1)
			idx++
		}
		for match != nil {
			ma(match)
			match, err = re.FindNextMatch(match)
			if err != nil {
				return nil, err
			}
		}
	}

	// 7、把所有的英文替换为行内语法格式
	re := regexp2.MustCompile(englishPattern, 0)
	match, err := re.FindStringMatch(fileData)
	if err != nil {
		return nil, err
	}

	seek := 0
	ma := func(match *regexp2.Match) {
		groups := match.Groups()
		raw := string(groups[0].Captures[0].Runes()) // 匹配到的字符串

		// 如果是纯数字，直接跳过
		_, err := strconv.Atoi(raw)
		if err == nil {
			return
		}

		if len(fileData) < seek {
			return
		}

		index := strings.Index(fileData[seek:], raw)
		done := fileData[:seek]
		handle := fileData[seek:]

		target := fmt.Sprintf("`%s`", raw)
		handle = strings.Replace(handle, raw, target, 1)
		fileData = done + handle
		seek += index + len(raw) + 2
	}
	for match != nil {
		ma(match)
		match, err = re.FindNextMatch(match)
		if err != nil {
			return nil, err
		}
	}

	// 把所有的标题去除特殊字符
	//// 用于把标题的反引号、加粗去除掉
	data, _ = RepairLevelHighLight([]byte(fileData))
	fileData = string(data)

	// 重新排序
	var orderMap map[string]string
	orderMap, data = OrderMapping(data, mMap)
	fileData = string(data)

	// 8、恢复图片、链接、代码块、行内语法格式

	seek = 0
	for i := 0; i < idx; i++ {
		if len(fileData) < seek {
			break
		}

		raw := fmt.Sprintf("@@%d@@", i)
		index := strings.Index(fileData[seek:], raw)
		done := fileData[:seek]
		handle := fileData[seek:]

		target := orderMap[raw]
		handle = strings.Replace(handle, raw, target, 1)
		fileData = done + handle
		seek += index
	}

	// 9、返回
	return []byte(fileData), nil
}
