package transport

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type PipeData struct {
	Value   string
	Handler func(string) string
	Next    chan string
}

func HandleClosedQueue() {
	queue := make(chan *PipeData, 10)
	out := make(chan string)
	go func() {
		queue <- &PipeData{
			Value:   "hello",
			Handler: DataCompute,
			Next:    out, //无缓冲下，接收发送配对
		}
		// 不关闭 queue 是无法 range 出数据的，会一直 blocking
		close(queue)
	}()

	go func() {
		for data := range queue {
			// 只有先发送才能关闭
			if data.Handler != nil && data.Value != "" {
				data.Next <- data.Handler(data.Value)
				close(data.Next)
			}
		}
	}()

	recv := <-out
	fmt.Println(recv)

}

const (
	GOLAYOUT = "2006-01-02 15:04:05"
)

// HandleQueue 无限循环，消费 N 条
func HandleQueue(size int) {
	queue := make(chan *PipeData, size)
	out := make(chan string)
	go func() {
		queue <- &PipeData{
			Value:   "hello",
			Handler: DataCompute,
			Next:    out,
		}
	}()

	go func() {
		for {
			data := <-queue
			if data.Handler != nil && data.Value != "" {
				data.Next <- data.Handler(data.Value)
				time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
				fmt.Println("[ data handled ] => ", data.Value, ",time=", time.Now().Format(GOLAYOUT))
			}
		}
	}()

	for {
		recv := <-out
		fmt.Println("[ queue consumed ] => ", recv, ",time=", time.Now().Format(GOLAYOUT))
		r := rand.Intn(100)
		queue <- &PipeData{
			Value:   "hello" + "-" + strconv.Itoa(r),
			Handler: DataCompute,
			Next:    out,
		}
		time.Sleep(time.Millisecond * 3000)
		fmt.Println("[ queue polling ] => ", time.Now().Format(GOLAYOUT))
	}

}

func DataCompute(s string) string {
	return strings.Map(func(r rune) rune {
		return unicode.ToUpper(r)
	}, s)
}
