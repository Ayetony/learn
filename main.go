package main

import (
	"fmt"

	"example.com/hello/heartbeat"
	"example.com/hello/transport"
	"example.com/hello/utils"
	"github.com/fatih/color"
)

func main() {
	// testing block
	transport.NoCSPVailable()
	color.Blue("Hello, Go Modules!")
	fmt.Println("module: example.com/hello")
	transport.LoadContent("do not repeating too much")
	var i = 10
	chekced := utils.CheckValue(i)
	for range 3 {
		if chekced {
			heartbeat.Beat()
		}
	}

	fmt.Print("Hello" > "world")
}
