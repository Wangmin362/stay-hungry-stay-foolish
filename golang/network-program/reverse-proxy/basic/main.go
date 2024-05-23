package main

import (
	"golang.org/x/net/http/httpproxy"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

// 目标服务器的URL （真实服务的地址）
const targetURL = "http://172.30.3.236:8090" // 您可以替换为您希望代理到的目标服务器

func main() {
	// 将目标URL转换为URL对象
	target, err := url.Parse(targetURL)
	if err != nil {
		log.Fatalf("无法解析目标URL: %v", err)
	}

	hpc := httpproxy.Config{
		HTTPProxy:  "socks5://172.30.3.224:10805",
		HTTPSProxy: "socks5://172.30.3.224:10805",
	}

	// 创建一个反向代理
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.Transport = &http.Transport{
		Proxy: func(r *http.Request) (*url.URL, error) {
			return hpc.ProxyFunc()(r.URL)
		},
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	// 创建一个HTTP处理函数，将请求转发到目标服务器
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r) // 通过反向代理处理请求
	})

	// 启动HTTP服务器
	log.Println("反向代理服务器启动，监听在 :8080")
	if err = http.ListenAndServe(":8880", nil); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
