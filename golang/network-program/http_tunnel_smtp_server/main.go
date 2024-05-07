package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

func handleCONNECT(w http.ResponseWriter, r *http.Request) {
	// 解析目标地址和端口
	host := r.URL.Hostname()
	port := r.URL.Port()

	// 建立到目标地址和端口的连接
	targetConn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		http.Error(w, "Failed to connect to target server", http.StatusInternalServerError)
		log.Println("Failed to connect to target server:", err)
		return
	}
	defer targetConn.Close()

	// 发送连接已建立的响应
	w.WriteHeader(http.StatusOK)
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Failed to hijack connection", http.StatusInternalServerError)
		log.Println("Failed to hijack connection")
		return
	}
	clientConn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, "Failed to hijack connection", http.StatusInternalServerError)
		log.Println("Failed to hijack connection:", err)
		return
	}
	defer clientConn.Close()

	// 将客户端连接和目标连接进行直接连接
	go func() {
		_, err := io.Copy(clientConn, targetConn)
		if err != nil {
			log.Println("Error copying from target to client:", err)
		}
	}()

	_, err = io.Copy(targetConn, clientConn)
	if err != nil {
		log.Println("Error copying from client to target:", err)
	}
}

func main() {
	http.HandleFunc("/", handleCONNECT)
	fmt.Println("HTTP CONNECT proxy server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
