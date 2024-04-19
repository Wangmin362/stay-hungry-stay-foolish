package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	// 查找指定服务的SRV记录
	_, srvRecords, err := net.LookupSRV("xmpp-server", "tcp", "google.com")
	if err != nil {
		fmt.Println("SRV 记录查询失败:", err)
		return
	}

	for _, srv := range srvRecords {
		fmt.Printf("服务地址: %s:%d\n", srv.Target, srv.Port)
		fmt.Printf("优先级: %d, 权重: %d\n", srv.Priority, srv.Weight)
		fmt.Println("-------------------------------")
	}

	server := http.Server{}
	server.ListenAndServe()
}
