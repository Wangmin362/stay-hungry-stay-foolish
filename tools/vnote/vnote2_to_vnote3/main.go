package main

import (
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var rootPath string = "D:\\Notebook\\Vnote"

//var rootPath string = "D:\\TMP\\Vnote\\Docker-bak"

var Empty = struct{}{}

// 用于迁移从vnote2到vnote3的数据
func main() {
	if err := RenameImagePath(rootPath); err != nil {
		panic(err)
	}
}

// RenameImagePath 返回值key为图片名
func RenameImagePath(root string) error {
	err := filepath.Walk(root, func(currPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if root == currPath {
			return nil
		}
		if info.IsDir() {
			if info.Name() == "Image" {
				return nil
			}
			return RenameImagePath(currPath)
		}

		ext := filepath.Ext(currPath)
		if ext != ".md" {
			return nil
		}

		file, err := os.OpenFile(currPath, os.O_RDWR, os.ModePerm)
		if err != nil {
			return err
		}

		defer file.Close()
		data, err := io.ReadAll(file)
		if err != nil {
			return err
		}
		newData := strings.Replace(string(data), "_v_images", "vx_images/", -1)
		_, err = file.Seek(0, 0)
		if err != nil {
			return err
		}
		_, err = file.Write([]byte(newData))
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func MoveDirToDir(src, dst string) error {
	err := filepath.Walk(src, func(currPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if src == currPath {
			return nil
		}

		newPath := path.Join(dst, info.Name())
		if info.IsDir() { // 如果当前是目录，那么就在目标路径中创建一个新的目录，然后执行拷贝
			if err := os.Mkdir(newPath, os.ModePerm); err != nil {
				return err
			}
			return MoveDirToDir(currPath, newPath)
		}

		// 如果是文件，直接重命名
		if err = os.Rename(currPath, newPath); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	if err = os.Remove(src); err != nil {
		return err
	}

	return nil
}
