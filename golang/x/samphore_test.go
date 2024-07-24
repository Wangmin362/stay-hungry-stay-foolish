package x

import (
	"context"
	"golang.org/x/sync/semaphore"
	"runtime"
	"testing"
)

var (
	maxWorkers = runtime.GOMAXPROCS(0)
	sema       = semaphore.NewWeighted(int64(maxWorkers))
	task       = make([]int, maxWorkers*4)
)

func TestWeighted(t *testing.T) {
	ctx := context.Background()

	for i := range task {
		if err := sema.Acquire(ctx, 1); err != nil {
			break
		}

		go func(i int) {
			defer sema.Release(1)
			task[i] = i + 1
		}(i)
	}

	//通过获取全部数量的信号量，来保证所有的worker全部处理完
	if err := sema.Acquire(ctx, int64(maxWorkers)); err != nil {
		t.Error("获取所有的worker失败")
	}
	t.Log(task)
}
