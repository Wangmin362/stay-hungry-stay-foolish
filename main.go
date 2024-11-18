package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	vDir := "E:\\电视剧\\暗黑者1+2\\"
	filepath.Walk(vDir, func(path string, info fs.FileInfo, err error) error {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		if vDir == path {
			return nil
		}
		if !info.IsDir() {
			return nil
		}

		return filepath.Walk(path, func(p string, info fs.FileInfo, err error) error {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println(err)
				}
			}()
			if path == p {
				return nil
			}
			if info.IsDir() {
				return nil
			}
			newPath := filepath.Join(filepath.Dir(p), strings.ReplaceAll(filepath.Base(p), "pdf", "mp4"))
			fmt.Println(newPath)
			os.Rename(p, newPath)
			return nil
		})
	})
}
