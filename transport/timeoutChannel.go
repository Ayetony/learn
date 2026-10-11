package transport

import (
	"fmt"
	"time"
)

func TimeoutHandle(ch chan string) {

	timeout := make(chan bool, 1)
	// 匿名函数实现超时控制，其实就是 channel 固定周期给一个写入值
	go func() {
		time.Sleep(1000 * time.Millisecond)
		timeout <- true
	}()
	// 使用 select 机制，辅助以定时写入，避免永久等待
	select {
	case readVal := <-ch:
		fmt.Println(readVal)
	case t := <-timeout:
		fmt.Printf("Timeout : %v ,at %s", t, time.Now().Format(time.DateTime))
		// 一直没有从 ch 读到数据，但是可以从 timeout 读到，因为 timeout 有固定周期写入

	}
}

func NoCSPVailable() {
	// no any caching
	// channel 发送和接收无法和自己配对，没有缓冲和default处理下，则死锁
	// chan 同步通信：发送方和接收方必须“碰头”才能完成一次传递。
	// 无缓冲 channel 的发送和接收必须同时配对
	ch := make(chan int)
	select {
	case <-ch:
	case ch <- 1:
	}
	fmt.Println(len(ch))
}

func NoDeadLockChannelWithCache() {
	ch := make(chan int, 1)
	ch <- 1
	<-ch
}

func DeadLockWithoutCache() {
	ch := make(chan int)
	ch <- 1
	<-ch
}
