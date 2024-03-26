package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang/demo/tools/oss/aliyun/app/sync"
)

func init() {

	dir := fmt.Sprintf("ossSyncer")
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	file := fmt.Sprintf("./%s/%s.txt", dir, time.Now().Format("2006-01-02"))
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}

	log.SetOutput(logFile) // 将文件设置为log输出的文件
	//log.SetPrefix("[ossAliyun]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	return
}

// 1、本地文件删除之后暂时不考虑删除云端的文件，保留备份，以免后面还需要
// TODO 2、考虑目录的重命名暂时不处理，后续写一个定时任务，直接清楚阿里云OSS中没有使用的图片
// TODO 如何保证图片的安全？ 防止其他人胡乱使用？  1、设置Refer done
// TODO 清理本地没有引用的图片
// TODO 日志输出到文件
// TODO 后台进程，开机自启动
// go build -o D:\Software\AliOssSyncer\aliOssSyncer.exe .\tools\oss\aliyun\app\sync\cmd\
func main() {
	syncer, err := sync.NewSyncer()
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	for {
		syncer.Run()
		time.Sleep(10 * time.Minute)
	}
}
