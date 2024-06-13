package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 打开文件
	file, err := os.Open("C:\\Users\\David\\Desktop\\yingyu.txt")
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	// 创建一个 Scanner 对象
	scanner := bufio.NewScanner(file)

	// 逐行读取文件内容
	for scanner.Scan() {
		// 输出每一行的内容
		line := scanner.Text()
		if line == "" {
			continue
		}
		fmt.Println(line)
	}

	// 检查是否有错误发生
	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件时发生错误:", err)
	}
}
