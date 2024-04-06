package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 使用http.Get函数请求外部服务
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		fmt.Printf("请求失败: %s\n", err)
		return
	}
	defer resp.Body.Close() // 确保关闭响应体

	// 读取响应体内容
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %s\n", err)
		return
	}

	// 输出获取到的外网IP
	fmt.Printf("当前的外网IP是: %s\n", string(ip))
}
