package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var rootPath string = "D:\\Notebook\\Vnote"

var Empty = struct{}{}

func main() {
	ClearImage(rootPath)
}

// ClearImage 返回值key为图片名
func ClearImage(rootPath string) {
	images := make(map[string]string, 5000)
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.Contains(path, "\\Image\\") {
			//fmt.Println(rootPath)
			images[info.Name()] = path
		}
		return nil
	})

	for image, imagePath := range images {
		containerImage, err := ContainerImage(rootPath, image)
		if err != nil {
			panic(err)
		}
		if !containerImage {
			fmt.Println(image, imagePath)
			if err := os.Remove(imagePath); err != nil {
				panic(err)
			}
		}
	}

	if err != nil {
		panic(err)
	}
}

func ContainerImage(rootPath, image string) (bool, error) {
	isContainer := false
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		if !strings.Contains(path, "Image") && strings.Contains(info.Name(), ".md") {
			readFile, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			if strings.Contains(string(readFile), image) {
				isContainer = true
				return nil
			}
		}

		return nil
	})
	return isContainer, err
}
