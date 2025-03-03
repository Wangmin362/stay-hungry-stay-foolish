package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// 国内镜像源列表
var mirrorSources = []string{
	"registry.cn-hangzhou.aliyuncs.com", // 阿里云镜像源
	"mirror.baidubce.com",               // 百度云镜像源
	"docker.mirrors.ustc.edu.cn",        // 中科大镜像源
}

// 添加时间戳到日志信息
func addTimestamp(message string) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("[%s] %s", timestamp, message)
}

// 拉取镜像函数
func pullImage(ctx context.Context, cli *client.Client, imageName string, mirror string) error {
	// 构建带有镜像源的镜像名称
	mirroredImage := fmt.Sprintf("%s/%s", strings.TrimRight(mirror, "/"), strings.TrimLeft(imageName, "/"))
	fmt.Print(addTimestamp(fmt.Sprintf("[start] Pulling image from %s: %s\n", mirror, mirroredImage)))

	// 尝试拉取镜像
	out, err := cli.ImagePull(ctx, mirroredImage, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer out.Close()

	// 简单打印拉取信息
	fmt.Println(addTimestamp(fmt.Sprintf("[end] Pulling image from %s: %s\n", mirror, mirroredImage)))
	return nil
}

func main() {
	// 创建 Docker 客户端
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Failed to create Docker client: %v", err)
	}
	defer cli.Close()

	// 获取用户输入的镜像名称
	if len(os.Args) < 2 {
		log.Fatal("Please provide an image name as an argument.")
	}
	imageName := os.Args[1]

	ctx := context.Background()

	// 尝试从每个镜像源拉取镜像
	for _, mirror := range mirrorSources {
		fmt.Printf("Trying to pull image from %s...\n", mirror)
		err := pullImage(ctx, cli, imageName, mirror)
		if err == nil {
			fmt.Print(addTimestamp(fmt.Sprintf("Successfully pulled image %s from %s\n", imageName, mirror)))
			return
		}
		fmt.Printf("Failed to pull image from %s: %v\n", mirror, err)
	}

	log.Fatalf("Failed to pull image %s from all mirror sources.", imageName)
}
