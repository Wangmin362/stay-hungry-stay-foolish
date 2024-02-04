package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	qDir := []string{"E:\\电视剧\\鸡毛飞上天\\", "E:\\电视剧\\平凡的世界\\", "E:\\电视剧\\射雕英雄传\\", "E:\\电视剧\\守护丽人\\",
		"E:\\电视剧\\天龙八部\\", "E:\\电视剧\\天天有喜\\", "E:\\电视剧\\蜗居\\", "E:\\电视剧\\幸福还会来敲门\\",
		"E:\\电视剧\\一仆二主\\"}
	for _, wdir := range qDir {
		resizeVideo(wdir, "Very Fast 480p30", "26")
	}
}

// HandBrakeCLI -i 01.ts --preset "Very Fast 480p30" -q 26 --optimize -o 01_rc26.mp4
func resizeVideo(vDir, preset, quality string) {
	filepath.Walk(vDir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		newVideo := fmt.Sprintf("%s\\new-%s", filepath.Dir(path), info.Name())
		cmd := exec.Command("HandBrakeCLI", "-i", path, "--preset", preset, "-q", quality, "--optimize", "-o", newVideo)
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout // 标准输出
		cmd.Stderr = &stderr // 标准错误
		err = cmd.Run()
		outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
		fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		} else {
			if err := os.Remove(path); err != nil {
				fmt.Printf("删除原始视频文件错误：%+v\n", err)
			} else {
				if err := os.Rename(newVideo, path); err != nil {
					fmt.Printf("把%s重命名为%s错误：%+v\n", newVideo, path, err)
				}
			}
		}
		return nil
	})
}
