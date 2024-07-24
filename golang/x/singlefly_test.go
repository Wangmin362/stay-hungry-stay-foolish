package x

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"strconv"
	"sync"
	"testing"
	"time"
)

var tmpVar = "test_key"

func TestSinglefly(t *testing.T) {
	var g singleflight.Group
	var wg sync.WaitGroup

	// 创建10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//使用singleFlight的Do方法来进行并发控制
			ret, err, _ := g.Do(tmpVar, func() (interface{}, error) {
				//这里是模拟业务处理
				timeRet := getTime()
				return timeRet, nil
			})

			if err != nil {
				fmt.Println("报错了", err.Error())
			}

			fmt.Println(ret)
		}()
	}

	wg.Wait()
}

func getTime() string {
	fmt.Println("exec query...")
	return strconv.Itoa(int(time.Now().UnixNano()))
}
