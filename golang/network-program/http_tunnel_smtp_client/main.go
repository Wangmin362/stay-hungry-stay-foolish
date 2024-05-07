package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"net/textproto"
)

func main() {
	// 目标SMTP服务器地址和端口号
	smtpHost := "smtp.example.com"
	smtpPort := "25"
	proxyAddr := "proxy.example.com:8080" // 代理服务器地址和端口号

	// 连接代理服务器
	proxyConn, err := net.Dial("tcp", proxyAddr)
	if err != nil {
		fmt.Println("Failed to connect to the proxy server:", err)
		return
	}
	defer proxyConn.Close()

	// 发送CONNECT请求
	fmt.Fprintf(proxyConn, "CONNECT %s:%s HTTP/1.1\r\nHost: %s:%s\r\n\r\n", smtpHost, smtpPort, smtpHost, smtpPort)

	// 读取响应
	reader := bufio.NewReader(proxyConn)
	code, message, err := textproto.NewReader(reader).ReadResponse(200)
	if err != nil || code != 200 {
		fmt.Println("Failed to establish tunnel:", message, err)
		return
	}

	// 建立TLS连接
	tlsConn := tls.Client(proxyConn, &tls.Config{
		InsecureSkipVerify: true, // 忽略服务器证书验证（仅作示例，请不要在生产环境中使用）
	})
	defer tlsConn.Close()

	// 连接到SMTP服务器
	smtpClient := textproto.NewConn(tlsConn)
	defer smtpClient.Close()

	// 发送SMTP命令
	_, err = smtpClient.Cmd("EHLO example.com")
	if err != nil {
		fmt.Println("Failed to send EHLO command:", err)
		return
	}

	// 这里可以继续发送其他SMTP命令和处理响应...
}
