package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// 监听端口
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server is listening on localhost:8888")

	for {
		// 等待客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		// 处理客户端连接
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// 接收文件名和文件大小
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}
	fileInfo := string(buffer[:n])
	fmt.Println("Received file info:", fileInfo)

	// 打开文件准备接收数据
	file, err := os.Create("received_file.dat")
	if err != nil {
		fmt.Println("Error creating file:", err.Error())
		return
	}
	defer file.Close()

	// 接收文件数据
	receivedBytes, err := io.Copy(file, conn)
	if err != nil {
		fmt.Println("Error copying file:", err.Error())
		return
	}
	fmt.Printf("Received file successfully, %d bytes\n", receivedBytes)
}
