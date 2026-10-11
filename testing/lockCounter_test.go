package testing

import (
	"runtime"
	"sync"
	"testing"

	"example.com/hello/transport"
)

func TestCounterIncrement(t *testing.T) {
	lock := &sync.Mutex{}
	for range 10 {
		go transport.Count(lock)
	}

	for {
		lock.Lock()
		c := transport.Counter
		lock.Unlock()
		// 主动让出当前 goroutine cpu 执行的机会，让 go 调度器运行其他 goroutine，不会挂起，保持运行状态排队，之后自动恢复运行
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}

}
