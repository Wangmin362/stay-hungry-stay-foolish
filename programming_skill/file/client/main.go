package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// 连接服务器
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	// 打开文件准备发送数据
	file, err := os.Open("large_file.dat")
	if err != nil {
		fmt.Println("Error opening file:", err.Error())
		return
	}
	defer file.Close()

	// 发送文件数据
	sentBytes, err := io.Copy(conn, file)
	if err != nil {
		fmt.Println("Error sending file:", err.Error())
		return
	}
	fmt.Printf("Sent file successfully, %d bytes\n", sentBytes)
}
