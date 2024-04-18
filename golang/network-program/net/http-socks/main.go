package main

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/proxy"
)

func main() {
	// SOCKS5 代理地址
	proxyAddr := "localhost:1080"

	// 创建一个代理拨号器
	dialer, err := proxy.SOCKS5("tcp", proxyAddr, nil, proxy.Direct)
	if err != nil {
		log.Fatal("Error creating SOCKS5 dialer:", err)
	}

	// 使用代理拨号器设置 HTTP 传输层
	transport := &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return dialer.Dial(network, addr)
		},
	}

	// 创建 HTTP 客户端
	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}

	// 发送 HTTP GET 请求
	resp, err := client.Get("http://example.com")
	if err != nil {
		log.Fatal("Error making HTTP GET request:", err)
	}
	defer resp.Body.Close()

	// 输出响应状态
	log.Println("Response status:", resp.Status)

	// 读取并打印响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}
	log.Println("Response body:", string(body))
}
