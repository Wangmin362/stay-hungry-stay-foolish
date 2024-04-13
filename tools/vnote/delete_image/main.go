package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// 删除所有的图片 TODO 注意，需要保证同步阿里云的程序再运行，这个时候就可以保证本地的图片肯定都上传到了阿里云，否则图片可能会丢失
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

		if strings.Contains(path, "vx_images") {
			ext := filepath.Ext(path)
			ext = strings.ToLower(ext)
			switch ext {
			case ".jpg", ".jpeg", ".bmp", ".png", ".webp", ".jif", ".svg":
				if err = os.Remove(path); err != nil {
					fmt.Printf("删除%s失败: %s\n", path, err.Error())
				} else {
					fmt.Printf("删除%s成功\n", path)
				}
			default:
				// ignore
			}
		}
		return nil
	})
}
