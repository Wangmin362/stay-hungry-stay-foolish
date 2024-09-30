package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir := "E:\\TMP"
	md5cache := make(map[string]string)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() { // 跳过目录
			hash, err := md5File(path)
			if err != nil {
				return err
			}

			oldPath, ok := md5cache[hash]
			if ok {
				fmt.Printf("%s, %v, %v\n", hash, oldPath, path)

				if err := os.Remove(path); err != nil {
					fmt.Printf("delete %v file error: %v\n", path, err)
				} else {
					fmt.Printf("delete replicate %v file successfule\n", err)
				}
				return nil
			}

			md5cache[hash] = path
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func md5File(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:]), nil
}
