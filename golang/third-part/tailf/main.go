package main

import (
	"fmt"

	"github.com/hpcloud/tail"
)

// 用于监听文件的变化，类似于linux的tail -f功能  linux下测试问题不大
func main() {
	t, err := tail.TailFile("/opt/test/fswatch", tail.Config{Follow: true})
	if err != nil {
		panic(err)
	}
	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}
