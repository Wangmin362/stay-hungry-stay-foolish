package main

import (
	"fmt"
	"golang.org/x/net/http/httpproxy"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	hpc := httpproxy.Config{
		HTTPProxy:  "socks5://172.30.3.224:10805",
		HTTPSProxy: "socks5://172.30.3.224:10805",
	}

	clt := http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			// 配置环境变量 HTTP_PROXY=socks5://192.168.11.3.224:10801 支持SOCKS5代理
			// 配置环境变量 HTTP_PROXY=192.168.11.3.224:10801          支持HTTP代理
			// 配置环境变量 HTTPS_PROXY=socks5://192.168.11.224:10801  支持SOCKS5代理
			// 配置环境变量 HTTPS_PROXY=192.168.11.224:10801           支持HTTPS代理
			// 即可实现SOCKS代理
			//Proxy: http.ProxyFromEnvironment,
			Proxy: func(r *http.Request) (*url.URL, error) {
				return hpc.ProxyFunc()(r.URL)
			},
		},
	}
	resp, err := clt.Get("http://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(all))
}
