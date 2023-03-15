package main

import (
	"fmt"
	"github.com/hpcloud/tail"
)

// 用于监听文件的变化，类似于linux的tail -f功能
func main() {
	t, err := tail.TailFile("C:\\Users\\wangmin\\Desktop\\fswatcher.txt", tail.Config{Follow: true})
	if err != nil {
		panic(err)
	}
	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}
