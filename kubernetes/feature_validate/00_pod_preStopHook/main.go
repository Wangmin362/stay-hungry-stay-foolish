package main

import (
	"flag"
	"log"
	"net/http"
	"time"
)

func Shutdown(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Unix()
	log.Println("shutdown start, waiting!!!")
	time.Sleep(time.Duration(delay) * time.Second)
	// MTA的业务流程

	log.Printf("shutdown, done!!!, spend=%ds", time.Now().Unix()-now)
	w.Write([]byte("关闭MTA容器\n"))
}

var delay int

func main() {
	flag.IntVar(&delay, "delay", 5, "延时时常")
	flag.Parse()

	http.HandleFunc("/mta/shutdown", Shutdown) // 设置访问的路由
	err := http.ListenAndServe(":19090", nil)  // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
