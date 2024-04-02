package main

import (
	"fmt"
	"syscall"
)

// 使用syscall包创建文件并写入内容
func main() {
	// 定义要创建的文件路径和内容
	filePath := "example.txt"
	content := []byte("Hello, syscall!")

	// 使用syscall包打开文件，如果不存在则创建，返回文件描述符
	fd, err := syscall.Open(filePath, syscall.O_CREAT|syscall.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer syscall.Close(fd) // 确保文件描述符最后被关闭

	// 使用syscall包写入内容
	_, writeErr := syscall.Write(fd, content)
	if writeErr != nil {
		panic(writeErr)
	}

	fmt.Println("File written successfully")
}
