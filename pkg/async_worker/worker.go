package async_worker

import (
	"context"
	"log"
	"runtime/debug"
	"sync"
	"sync/atomic"
)

// AsyncWorker 简化版的协程池
// 适合处理无状态的任务，且无需感知任务结果，允许成功或失败，不设置过期时间
type AsyncWorker[T any] interface {
	Submit(req T)
	TrySubmit(req T) bool
	Wait()
	Stop()
}

type asyncWorker[T any] struct {
	ch chan T
	wg *sync.WaitGroup

	closed atomic.Int32
}

func NewAsyncWorker[T any](ctx context.Context, n int, taskFn func(ctx context.Context, req T)) AsyncWorker[T] {
	ch := make(chan T, n)
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		go func() {
			defer func() {
				if r := recover(); r != nil {
					log.Fatalf("[panic] catch: %v\n%s", r, debug.Stack())
				}
			}()
			for item := range ch {
				taskFn(ctx, item)
				wg.Done()
			}
		}()
	}
	return &asyncWorker[T]{
		ch: ch, wg: &wg,
	}
}

func (aw *asyncWorker[T]) Submit(req T) {
	if aw.closed.Load() == 0 {
		aw.wg.Add(1)
		aw.ch <- req
	}
}

func (aw *asyncWorker[T]) TrySubmit(req T) bool {
	if aw.closed.Load() == 0 {
		aw.wg.Add(1)
		select {
		case aw.ch <- req:
			return true
		default:
			aw.wg.Done()
			return false
		}
	}
	return false
}

func (aw *asyncWorker[T]) Wait() {
	aw.wg.Wait()
}

func (aw *asyncWorker[T]) Stop() {
	aw.wg.Wait()
	if success := aw.closed.CompareAndSwap(0, 1); success {
		close(aw.ch)
	}
}
