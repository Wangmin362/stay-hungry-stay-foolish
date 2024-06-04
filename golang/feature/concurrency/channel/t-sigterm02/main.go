package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	ctx22, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	ctx333, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			<-ctx.Done()
			fmt.Println("stopping")
			return
		}
	}()
	go func() {
		for {
			<-ctx22.Done()
			fmt.Println("stopping222")
			return
		}
	}()
	go func() {
		for {
			<-ctx333.Done()
			fmt.Println("stopping333")
			return
		}
	}()

	fmt.Println("running")
	time.Sleep(10 * time.Hour)
}
