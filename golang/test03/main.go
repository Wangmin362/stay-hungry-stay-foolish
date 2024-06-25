package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// 读取待压缩的文件内容
	fileContent, err := os.Open("D:\\Skyguard\\discovery-2024-06-05-15-03-57")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 构建请求
	url := "https://bj-ucss-230.gatorcloud.skyguardmis.com/skgwSps/sps/v1/ztna/privateApp/discovery/discovery-2024-06-05-15-03-57"
	req, err := http.NewRequest(http.MethodPut, url, fileContent)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Encoding", "gzip")
	req.Header.Set("content-type", "application/octet-stream")
	req.Header.Set("authorization", "Basic Mjk2MjE3NjU0MTo4NmQ0ZDRhZDc0ZmE3YWQ2ZjQxOTk3YzRlZWM3N2IxZmJmM2YyZmViYzBkZjM0YjYyZWQ5YzMyMDUyMDIxZGE0OjEwMDAwMDU6Njg5ZjkwNTctZGQyMS00ZTI0LTk0ZTItOGUyYWRiMTBiZThh")
	req.Header.Set("host", "bj-ucss-230.gatorcloud.skyguardmis.com")
	req.Header.Set("x-pop-id", "689f9057-dd21-4e24-94e2-8e2adb10be8a")
	req.Header.Set("x-tenant-id", "1000005")
	req.Header.Set("x-timestamp", "2962176541")

	// 发送请求
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 打印响应
	fmt.Println("Response status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response body:", string(body))
}
