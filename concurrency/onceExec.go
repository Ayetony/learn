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
