package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	for {
		resizeVideo("E:\\电视剧\\", "Very Fast 480p30", "26")
		time.Sleep(1 * time.Minute)
	}
}

// HandBrakeCLI -i 01.ts --preset "Very Fast 480p30" -q 26 --optimize -o 01_rc26.mp4
func resizeVideo(vDir, preset, quality string) {
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
			resizeVideo(path, preset, quality)
		}

		optSize := int64(300 * 1024 * 1024)
		if info.Size() < optSize {
			return nil
		}

		ext := filepath.Ext(info.Name())
		switch ext {
		case ".part", ".downloading", ".xltd":
			fmt.Println(path, ", 正在下载！！！")
			return nil
		}

		start := time.Now().Unix()
		fmt.Println(fmt.Sprintf("【%s】开始处理视频：%s", time.Now().Format("2006-01-02 15:04:05"), path))
		newVideo := fmt.Sprintf("%s\\new-%s", filepath.Dir(path), info.Name())
		cmd := exec.Command("HandBrakeCLI", "-i", path, "--preset", preset, "-q", quality, "--optimize", "-o", newVideo)
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
				if err := os.Rename(newVideo, path); err != nil {
					fmt.Printf("把%s重命名为%s错误：%+v\n", newVideo, path, err)
				}
			}
		}
		total := (time.Now().Unix() - start) / 60
		fmt.Println(fmt.Sprintf("【%s】耗时%d分钟，处理完成视频：%s", time.Now().Format("2006-01-02 15:04:05"), total, path))
		return nil
	})
}
