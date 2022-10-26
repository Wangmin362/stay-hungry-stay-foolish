package sync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// TestOnce1 Once保证需要运行的函数只被运行一次
func TestOnce1(t *testing.T) {
	one := sync.Once{}

	for {
		// 无论这个for循环执行多少次，once只会被运行一次
		one.Do(func() {
			fmt.Println("once")
		})
		time.Sleep(time.Second)
	}
}

// TestOnce2 错误的使用方式，多次实例化Once
func TestOnce2(t *testing.T) {
	for {
		// 每次循环都会实例化一个Once,因此每次循环都会执行一次函数
		one := sync.Once{}
		one.Do(func() {
			fmt.Println("once")
		})
		time.Sleep(time.Second)
	}
}
