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
	flacToMp3("E:\\")
}

// ffmpeg -i input.flac -ab 320k -map_metadata 0 -id3v2_version 3 output.mp3
func flacToMp3(vDir string) {
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

		ext := filepath.Ext(info.Name())
		if ext != ".flac" {
			return nil
		}

		start := time.Now().UnixMilli()
		fmt.Println(fmt.Sprintf("【%s】开始处理音频：%s", time.Now().Format("2006-01-02 15:04:05"), path))
		newVideo := fmt.Sprintf("%s\\%s", filepath.Dir(path), strings.ReplaceAll(info.Name(), "flac", "mp3"))
		cmd := exec.Command("ffmpeg", "-i", path, "-ab", "320k", "-map_metadata", "0", "-id3v2_version", "3", newVideo)
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout // 标准输出
		cmd.Stderr = &stderr // 标准错误
		err = cmd.Run()
		if err != nil {
			log.Printf("cmd.Run() failed with %s\n", err)
			return nil
		} else {
			if err := os.Remove(path); err != nil {
				fmt.Printf("删除原始音频文件错误：%+v\n", err)
			}
		}
		total := time.Now().UnixMilli() - start
		fmt.Println(fmt.Sprintf("【%s】耗时%d毫秒，处理完成音频：%s", time.Now().Format("2006-01-02 15:04:05"), total, path))
		fmt.Println()
		return nil
	})
}
