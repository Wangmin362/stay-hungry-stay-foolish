package main

import (
	"context"
	"fmt"
	"github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
	"log"
)

func main() {
	// 连接到默认的containerd套接字
	client, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		log.Fatalf("Failed to connect to containerd: %v", err)
	}
	defer client.Close()

	// 创建一个新的命名空间
	ctx := namespaces.WithNamespace(context.Background(), "k8s.io")

	// 获取容器列表
	containers, err := client.Containers(ctx)
	if err != nil {
		log.Fatalf("Failed to get container list: %v", err)
	}

	// 遍历容器列表并输出名称
	for _, container := range containers {
		info, err := container.Info(ctx)
		if err != nil {
			log.Printf("Failed to get container info: %v", err)
			continue
		}
		fmt.Printf("%+v\n", info)
	}
}
