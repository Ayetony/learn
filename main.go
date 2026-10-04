package main

import (
	"fmt"

	"com.practice/tools"
	"com.practice/transport"
	checkutils "example.com/hello/check-utils"
	"example.com/hello/heartbeat"
	"github.com/fatih/color"
)

func main() {
	color.Green("Hello, Go Modules!")
	fmt.Println("module: example.com/hello")
	tools.RunDemo()
	transport.LoadContent("do not repeating too much")
	var i int = 10
	checkutils.CheckValue(i + 1)
	for range 5 {
		heartbeat.Beat()
	}
}
