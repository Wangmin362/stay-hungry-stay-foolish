package main

import (
	"fmt"
	"github.com/justinas/alice"
	"net/http"
)

func main() {
	// 构造一个中间件链
	myChain := alice.New(loggingMiddleware, recoveringMiddleware, authenticatingMiddleware)

	// 应用到一个HTTP处理函数
	http.Handle("/", myChain.Then(http.HandlerFunc(myAppHandler)))
	http.ListenAndServe(":8080", nil)
}

// 日志记录中间件
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 这里可以添加日志逻辑
		fmt.Println("日志记录...")
		next.ServeHTTP(w, r)
	})
}

// 错误恢复中间件
func recoveringMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 添加错误恢复逻辑
		fmt.Println("错误恢复...")
		next.ServeHTTP(w, r)
	})
}

// 认证中间件
func authenticatingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 进行认证检查
		fmt.Println("认证逻辑...")
		next.ServeHTTP(w, r)
	})
}

// 应用的主处理函数
func myAppHandler(w http.ResponseWriter, r *http.Request) {
	// 应用逻辑
	fmt.Println("业务逻辑...")
	w.Write([]byte("Hello, world! 🌍"))
}
