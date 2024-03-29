package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	goRun "runtime"
	"testing"
	"time"
)

func TestJetter(t *testing.T) {
	stop := make(chan struct{})
	old := time.Now().UnixMilli()
	wait.JitterUntil(func() {
		new := time.Now().UnixMilli()
		diff := new - old
		fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"), diff)
		old = new
	}, 5*time.Second, 0.5, true, stop)
}

func TestJetterContext(t *testing.T) {
	old := time.Now().UnixMilli()
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"))
	wait.JitterUntilWithContext(context.Background(), func(ctx context.Context) {
		new := time.Now().UnixMilli()
		diff := new - old
		fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"), diff)
		old = new
	}, 5*time.Second, 0.5, true)
}

func TestUntilWithContext(t *testing.T) {

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("协程的数量为：", goRun.NumGoroutine())
			time.Sleep(2 * time.Second)
		}
	}()

	ctx := context.Background()
	old := time.Now().UnixMilli()
	wait.UntilWithContext(ctx, func(ctx context.Context) {
		new := time.Now().UnixMilli()
		diff := new - old
		fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"), diff)
		old = new
		time.Sleep(time.Hour)
	}, time.Second*2)
}

func init() {
	runtime.ReallyCrash = false
}

func init() {
	fmt.Println("sssdss")
}

func TestUntilWithContextPanic(t *testing.T) {
	ctx := context.Background()
	old := time.Now().UnixMilli()
	wait.UntilWithContext(ctx, func(ctx context.Context) {
		new := time.Now().UnixMilli()
		diff := new - old
		fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"), diff)
		old = new
		panic(errors.Errorf("create panic"))
	}, time.Second*5)
}
