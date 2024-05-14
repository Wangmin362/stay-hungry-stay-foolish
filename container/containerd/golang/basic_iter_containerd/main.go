package main

import (
	"context"
	"fmt"
	"log"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/cio"
	"github.com/containerd/containerd/namespaces"
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
		label, err := container.Labels(ctx)
		if err != nil {
			log.Printf("Failed to get container info: %v", err)
			continue
		}

		kind := label["io.cri-containerd.kind"]
		podName := label["io.kubernetes.pod.name"]
		containerName := label["io.kubernetes.container.name"]

		if kind == "sandbox" {
			continue
		}

		spec, err := container.Spec(ctx)
		if err != nil {
			log.Printf("Failed to get container spec: %v", err)
			continue
		}

		task, err := container.Task(ctx, cio.NewAttach())
		if err != nil {
			continue
		}
		pid := task.Pid()

		fmt.Printf("%s -> %s, pid=%d\n", podName, containerName, pid)

		for _, mount := range spec.Mounts {
			fmt.Printf("type=%s, src=%s, dst=%s \n", mount.Type, mount.Source, mount.Destination)
		}

		fmt.Println()

	}
}
