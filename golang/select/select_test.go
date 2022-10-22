package main

import (
	"fmt"
	"testing"
	"time"
)

// TestSelect1 select中的break只能跳出select，如果select外面报了一层for循环,是无法直接退出的
func TestSelect1(t *testing.T) {
	cnt := 1
	for {
		select {
		default:
			fmt.Println("select before break")
			time.Sleep(time.Second)
			cnt++
			if cnt == 5 {
				// 即便在select中break, 只能够阻断select后面的流程，并不能跳出外层的for循环
				fmt.Println("break")
				break
			}
			fmt.Println("select after break")

		}
	}
}
