package main

import (
	"crypto/tls"
	"log"
	"net/http"
)

func main() {
	// 配置多个证书
	cert1, err := tls.LoadX509KeyPair("cert1.crt", "cert1.key") // 替换为你的证书文件路径
	if err != nil {
		log.Fatal(err)
	}

	cert2, err := tls.LoadX509KeyPair("cert2.crt", "cert2.key") // 替换为你的证书文件路径
	if err != nil {
		log.Fatal(err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert1, cert2},
	}

	// 创建HTTPS服务器
	server := &http.Server{
		Addr:      ":8056",
		TLSConfig: tlsConfig,
	}

	// 处理请求
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, HTTPS World!"))
	})

	log.Println("Server started at https://localhost:8056")

	// 启动服务器
	err = server.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}
