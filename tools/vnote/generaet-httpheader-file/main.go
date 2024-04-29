package main

import (
	"fmt"
	"github.com/golang/demo/tools/vnote"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	dir := "D:/Notebook/Vnote/Blog/02.Linux网络/04.应用层/00.HTTP/05.Header"

	err := GenerateFile(dir)
	if err != nil {
		log.Fatal(err)
	}
}

var pattern []string = []string{
	"HTTP %s Header是什么？",
	"为什么需要HTTP %s Header？",
	"什么场景下适合HTTP %s Header？",
	"HTTP %s Header常见用法",
	"浏览器是如何处理HTTP %s Header的？",
	"使用HTTP %s Header注意事项",
	"使用Golang演示HTTP %s Header用法",
}

func GenerateFile(dir string) error {
	dirs := make(map[string]os.FileInfo, 1024)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if dir == path {
			return nil
		}

		if !info.IsDir() {
			return nil // 文件直接跳过
		}

		if strings.Contains(path, "_v_attachments") {
			return nil
		}

		split := strings.Split(info.Name(), ".")
		num, err := strconv.Atoi(split[0])
		if err != nil {
			return err
		}
		if num < 23 {
			return nil
		}

		dirs[path] = info

		return nil
	})
	if err != nil {
		return err
	}

	for tmpPath := range dirs {
		err := filepath.Walk(tmpPath, func(path string, info os.FileInfo, err error) error {
			if tmpPath == path {
				return nil
			}

			if info.IsDir() {
				return nil // 文件夹直接跳过
			}

			if err = os.Remove(path); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	//pp := "D:\\Notebook\\Vnote\\Blog\\02.Linux网络\\04.应用层\\00.HTTP\\05.Header\\23.Link"
	//in := dirs[pp]

	for dpath, dinfo := range dirs {
		split := strings.Split(dinfo.Name(), ".")
		dName := split[1]

		childVX, err := vnote.UnMarshal("D:/Notebook/Vnote/Blog/02.Linux网络/04.应用层/00.HTTP/05.Header/22.Last-Modified/vx.json")
		if err != nil {
			return err
		}
		childVX.Folders = []*vnote.Folder{}
		childVX.Files = []*vnote.File{}

		// 目录下下放入QA，以及各个问题，并修改上面的vx.json文件
		ptnQA, err := os.ReadFile("D:/Notebook/Vnote/Blog/02.Linux网络/04.应用层/00.HTTP/05.Header/22.Last-Modified/00.QA.md")
		if err != nil {
			return nil
		}
		ptnQA = []byte(strings.ReplaceAll(string(ptnQA), "Last-Modified", dName))

		childVX.Files = append(childVX.Files, &vnote.File{
			CreatedTime:  "2024-04-18T07:24:10Z",
			Id:           "6076",
			ModifiedTime: "2024-04-18T07:24:10Z",
			Name:         "00.QA.md",
			Signature:    "177807976771",
		})

		for idx, q := range pattern {
			fName := fmt.Sprintf("%02d.%s.md", idx+1, fmt.Sprintf(q, dName))
			qPath := filepath.Join(dpath, fName)
			if err = os.WriteFile(qPath, []byte{}, os.ModePerm); err != nil {
				return err
			}
			childVX.Files = append(childVX.Files, &vnote.File{
				CreatedTime:  "2024-04-18T07:24:10Z",
				Id:           "6076",
				ModifiedTime: "2024-04-18T07:24:10Z",
				Name:         fName,
				Signature:    "177807976771",
			})

			appQ := fmt.Sprintf("\n你是一个高级网络专家，请写一篇文章，详细阐述%s 请在文章中尽可能多的帮我添加emoji以增强文章趣味\n\n", fmt.Sprintf(q, dName))
			ptnQA = append(ptnQA, []byte(appQ)...)
		}

		if err = os.WriteFile(filepath.Join(dpath, "00.QA.md"), ptnQA, os.ModePerm); err != nil {
			return err
		}

		childVXPath := filepath.Join(dpath, "vx.json")
		if err = childVX.PersistentMarshal(childVXPath); err != nil {
			return err
		}
	}

	return nil
}
