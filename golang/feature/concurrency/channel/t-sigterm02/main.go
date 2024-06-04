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

	go func() {
		for {
			<-ctx.Done()
			fmt.Println("stopping")
			return
		}
	}()

	fmt.Println("running")
	time.Sleep(10 * time.Hour)
}
