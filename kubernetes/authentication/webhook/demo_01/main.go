package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var port string

// 参考文章：https://www.51cto.com/article/694171.html
func main() {
	flag.StringVar(&port, "port", "9999", "http  server  port")
	flag.Parse()
	//  启动httpserver
	wbsrv := WebHookServer{server: &http.Server{
		Addr: fmt.Sprintf(":%v", port),
	}}
	mux := http.NewServeMux()
	// TODO 这里的/auth URL肯定是可以进行配置的，K8S设计者不可能是写死的，估计是在某个配置文件中指定的
	mux.HandleFunc("/auth", wbsrv.serve)
	wbsrv.server.Handler = mux

	//  启动协程来处理
	go func() {
		if err := wbsrv.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			glog.Errorf("Failed  to  listen  and  serve  webhook  server:  %v", err)
		}
	}()

	glog.Info("Server  started")

	//  优雅退出
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	glog.Infof("Got  OS  shutdown  signal,  shutting  down  webhook  server  gracefully...")
	_ = wbsrv.server.Shutdown(context.Background())
}
