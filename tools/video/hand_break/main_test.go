package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"path"
	"testing"
	"time"
)

func TestPreset(t *testing.T) {

	presets := []string{"Very Fast 480p30", "Fast 480p30"}
	qualitys := []string{"18", "20", "22", "24", "26", "28", "30", "32", "34", "36"}

	dir := "C:\\Users\\wangmin\\Downloads"
	name := "速度与激情7.mp4"

	for _, preset := range presets {
		for _, quality := range qualitys {
			oldVideo := path.Join(dir, name)
			newVideo := path.Join(dir, fmt.Sprintf("%s_%s_%s", preset, quality, name))
			start := time.Now().Unix()
			cmd := exec.Command("HandBrakeCLI", "-i", oldVideo, "--preset", preset, "--encoder-preset", "vcn", "-q", quality, "--optimize", "-o", newVideo)
			var stdout, stderr bytes.Buffer
			cmd.Stdout = &stdout // 标准输出
			cmd.Stderr = &stderr // 标准错误
			err := cmd.Run()
			if err != nil {
				fmt.Printf("cmd.Run() failed with %s\n", err)
			}
			total := time.Now().Unix() - start
			fmt.Println(fmt.Sprintf("处理%s共消耗%d秒！！！", newVideo, total))
		}
	}
}
