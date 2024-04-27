package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my Go HTTP server!")
}

func main() {
	http.HandleFunc("/", handler) // 设置根目录的处理器
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil) // 在端口8080上启动服务器
}
