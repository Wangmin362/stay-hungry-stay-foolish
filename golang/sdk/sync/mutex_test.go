package sync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 测试go语言的零值可用
func TestZeroUseable(t *testing.T) {
	// lock变量没有进行初始化
	// todo var lock sync.Mutex 是否等价于 lock := sync.Mutext{}
	//var lock sync.Mutex
	lock := sync.Mutex{}
	fmt.Printf("addr is :%p", &lock)
	lock.Lock()
	time.Sleep(5 * time.Second)
	lock.Unlock()
}
