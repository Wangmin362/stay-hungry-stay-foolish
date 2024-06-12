package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var stopLock sync.Mutex
	stopChan := make(chan struct{}, 1)
	signalChan := make(chan os.Signal, 1)
	go func() {
		//阻塞程序运行，直到收到终止的信号
		<-signalChan
		stopLock.Lock()
		stopLock.Unlock()
		// 停止调试的时候是可以收到停止信号的
		log.Println("Cleaning before stop...")
		stopChan <- struct{}{}
		os.Exit(0)
	}()

	fmt.Println("print")
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	//模拟一个持续运行的进程
	time.Sleep(10 * time.Hour)
}
