package testing

import (
	"testing"

	"example.com/hello/transport"
)

func TestChanDeadLockWithCache(t *testing.T) {
	// 有缓冲下没有同步语配对不会死锁
	transport.NoDeadLockChannelWithCache()
}

func TestChanDeadLockWithoutCache(t *testing.T) {
	// 无缓冲的同步语义必须接受发送配对，否则死锁
	transport.DeadLockWithoutCache()
}
