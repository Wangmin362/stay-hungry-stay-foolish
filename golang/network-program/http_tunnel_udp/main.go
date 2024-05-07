package main

import (
	"github.com/elazarl/goproxy"
	"log"
	"net"
	"net/http"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)

	// 拦截 CONNECT 请求，然后建立 UDP 连接并进行代理
	proxy.OnRequest().DoFunc(
		func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			// 建立到目标 UDP 服务器的连接
			targetAddr := "udp.example.com:1234"
			conn, err := net.Dial("udp", targetAddr)
			if err != nil {
				log.Printf("Failed to connect to UDP server: %s", err)
				return nil, goproxy.NewResponse(req, goproxy.ContentTypeText, http.StatusServiceUnavailable, "Failed to connect to UDP server")
			}
			defer conn.Close()

			// 将请求的 CONNECT 数据发送到 UDP 服务器
			_, err = conn.Write([]byte("Hello, UDP server!"))
			if err != nil {
				log.Printf("Failed to send data to UDP server: %s", err)
				return nil, goproxy.NewResponse(req, goproxy.ContentTypeText, http.StatusInternalServerError, "Failed to send data to UDP server")
			}

			// 读取 UDP 服务器的响应数据
			buffer := make([]byte, 1024)
			_, err = conn.Read(buffer)
			if err != nil {
				log.Printf("Failed to read data from UDP server: %s", err)
				return nil, goproxy.NewResponse(req, goproxy.ContentTypeText, http.StatusInternalServerError, "Failed to read data from UDP server")
			}

			// 创建 HTTP 响应并返回
			return nil, goproxy.NewResponse(req, goproxy.ContentTypeText, http.StatusOK, string(buffer))
		},
	)

	// 启动代理服务器
	log.Fatal(http.ListenAndServe(":8080", proxy))
}
