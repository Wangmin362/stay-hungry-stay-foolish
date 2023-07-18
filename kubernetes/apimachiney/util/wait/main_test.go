package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/wait"
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