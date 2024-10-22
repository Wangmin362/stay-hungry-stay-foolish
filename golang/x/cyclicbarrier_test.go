package x

import (
	"fmt"
	"github.com/marusama/cyclicbarrier"
	"sync"
	"testing"
	"time"
)

func TestWaitCyclicBarrier(t *testing.T) {
	cnt := 0
	cb := cyclicbarrier.NewWithAction(10, func() error {
		cnt++
		return nil
	})
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 5; j++ {
				time.Sleep(100 * time.Millisecond)
				cb.Await(nil)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(cnt)
}
