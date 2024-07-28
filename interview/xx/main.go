package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// 启动TCP服务端
	go startTCPServer()

	// 模拟TCP客户端发送数据
	sendDataViaTCP()

	// 接收并输出缓存的数据
	readBufferedData()
}

func startTCPServer() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer l.Close()
	fmt.Println("TCP server listening on :8888")
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			return
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		// Read the incoming connection into the buffer.
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading:", err.Error())
			}
			break
		}
		// Print the received data.
		fmt.Printf("Received TCP data: %s\n", string(buf[:n]))

		// Here, you can cache the received data into a buffer.
		// For simplicity, let's cache it into a bytes.Buffer.
		buffer.Write(buf[:n])
	}
}

var buffer bytes.Buffer

func sendDataViaTCP() {
	// Connect to the TCP server.
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	// Send some data to the server.
	message := "Hello, TCP server!"
	fmt.Println("Sending TCP data:", message)
	conn.Write([]byte(message))
}

func readBufferedData() {
	// Now, you can read the buffered data using io.Reader interface.
	// For example, let's read and print the buffered data.
	fmt.Println("Reading buffered data:")
	io.Copy(os.Stdout, &buffer)
}
