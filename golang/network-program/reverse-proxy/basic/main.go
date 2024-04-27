package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// 目标服务器的URL
const targetURL = "http://baidu.com" // 您可以替换为您希望代理到的目标服务器

func main() {
	// 将目标URL转换为URL对象
	target, err := url.Parse(targetURL)
	if err != nil {
		log.Fatalf("无法解析目标URL: %v", err)
	}

	// 创建一个反向代理
	proxy := httputil.NewSingleHostReverseProxy(target)

	// 创建一个HTTP处理函数，将请求转发到目标服务器
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r) // 通过反向代理处理请求
	})

	// 启动HTTP服务器
	log.Println("反向代理服务器启动，监听在 :8080")
	if err = http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
