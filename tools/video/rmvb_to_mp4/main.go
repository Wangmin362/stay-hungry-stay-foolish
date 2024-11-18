package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	rmvbToMp4("G:\\test")
}

// HandBrakeCLI -i input.rmvb -o output.mp4 --preset="Fast 1080p30"
// ffmpeg -i [inputname].rmvb -c:v libx264 -strict -2 [outputname].mp4
// ffmpeg -i [inputname].rmvb -c:v libx264 -c:a aac -strict -2 [outputname].mp4
func rmvbToMp4(vDir string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	filepath.Walk(vDir, func(path string, info fs.FileInfo, err error) error {
		if vDir == path {
			return nil
		}
		if info.IsDir() {
			return nil
		}

		if strings.ToLower(filepath.Ext(path)) != ".rmvb" {
			return nil
		}

		start := time.Now().Unix()
		fmt.Println(fmt.Sprintf("【%s】开始处理视频：%s", time.Now().Format("2006-01-02 15:04:05"), path))
		name := info.Name()
		name = strings.ReplaceAll(name, ".rmvb", ".mp4")
		newVideo := fmt.Sprintf("%s\\new-%s", filepath.Dir(path), name)
		cmd := exec.Command("ffmpeg", "-i", path, "-c:v", "libx264", "-strict", "-2", newVideo)
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout // 标准输出
		cmd.Stderr = &stderr // 标准错误
		err = cmd.Run()
		if err != nil {
			log.Printf("cmd.Run() failed with %s\n", err)
			return nil
		} else {
			if err := os.Remove(path); err != nil {
				fmt.Printf("删除原始视频文件错误：%+v\n", err)
			} else {
				newpath := fmt.Sprintf("%s\\%s", filepath.Dir(path), name)
				if err := os.Rename(newVideo, newpath); err != nil {
					fmt.Printf("把%s重命名为%s错误：%+v\n", newVideo, path, err)
				}
			}
		}
		total := (time.Now().Unix() - start) / 60
		fmt.Println(fmt.Sprintf("【%s】耗时%d分钟，处理完成视频：%s", time.Now().Format("2006-01-02 15:04:05"), total, path))
		return nil
	})
}
