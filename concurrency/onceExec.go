// Package concurrency
package concurrency

import (
	"fmt"
	"sync"
)

var first string
var once sync.Once

func Setup() {
	first = "hello world!"
	// once 类型保证全局唯一操作，也就是只需要执行一次
	fmt.Println("Execute first only")
}

func DoPrint() {
	once.Do(Setup)
	println(first)
}

func TwoPrint() {
	go DoPrint()
	go DoPrint()
}
