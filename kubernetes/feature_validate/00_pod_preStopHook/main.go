package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func Shutdown(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Unix()
	log.Println("shutdown start, waiting!!!")
	s := time.Now().Format(time.DateTime)
	testFile, err := os.OpenFile("/test/abc.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0o666)
	if err != nil {
		log.Printf("/test/abc.txt文件打开错误:%+v\n", err)
	}
	defer testFile.Close()
	if _, err := testFile.WriteString(fmt.Sprintf("%s\n", s)); err != nil {
		log.Println("写入文件错误")
	}
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
