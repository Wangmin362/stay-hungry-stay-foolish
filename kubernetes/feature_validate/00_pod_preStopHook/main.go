package main

import (
	"log"
	"net/http"
)

func Shutdown(w http.ResponseWriter, r *http.Request) {
	log.Println("shutdown, waiting!!!")
	w.Write([]byte("关闭MTA容器"))
}
func main() {
	http.HandleFunc("/mta/shutdown", Shutdown) //设置访问的路由
	err := http.ListenAndServe(":19090", nil)  //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
