package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	clt := http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			// 配置环境变量 HTTP_PROXY=socks5://192.168.11.3.224:10801 支持HTTP SOCKS5代理
			// 配置环境变量 HTTPS_PROXY=socks5://192.168.11.224:10801  支持HTTPS SOCKS5代理
			// 即可实现SOCKS代理
			Proxy: http.ProxyFromEnvironment,
		},
	}
	resp, err := clt.Get("https://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(all))
}
