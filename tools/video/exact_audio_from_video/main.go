package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	exactAudio("G:\\")
}

// ffmpeg -i input.mp4 -f mp3 -vn output.mp3
// ffmpeg -i input.mp4 -vn -acodec libmp3lame -ab 192k output.mp3
func exactAudio(vDir string) {
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println(err)
	//	}
	//}()

	filepath.Walk(vDir, func(path string, info fs.FileInfo, err error) error {
		if vDir == path {
			return nil
		}
		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(info.Name())
		if ext != ".mp4" {
			fmt.Printf("当前%s文件不是mp4格式\n", path)
			return nil
		}

		start := time.Now().UnixMilli()
		fmt.Println(fmt.Sprintf("【%s】开始处理视频：%s", time.Now().Format("2006-01-02 15:04:05"), path))
		newVideo := fmt.Sprintf("%s\\%s", filepath.Dir(path), strings.ReplaceAll(info.Name(), "mp4", "mp3"))
		newVideo = strings.ReplaceAll(newVideo, "G:\\04_音乐mv\\", "E:\\音乐\\")
		//cmd := exec.Command("ffmpeg", "-i", path, "-f", "mp3", "-vn", newVideo)
		cmd := exec.Command("ffmpeg", "-i", path, "-vn", "-acodec", "libmp3lame", "-ab", "192k", newVideo)
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout // 标准输出
		cmd.Stderr = &stderr // 标准错误
		err = cmd.Run()
		if err != nil {
			log.Printf("cmd.Run() failed with %s\n", err)
			return nil
		}
		total := time.Now().UnixMilli() - start
		fmt.Println(fmt.Sprintf("【%s】耗时%d毫秒，处理完成视频：%s", time.Now().Format("2006-01-02 15:04:05"), total, path))
		return nil
	})
}
