package main

import (
	"fmt"

	"example.com/hello/transport"
	"github.com/fatih/color"
)

func main() {
	// testing block
	transport.HandleQueue(3)
	transport.CSPVailable()
	color.Blue("Hello, Go Modules!")
	fmt.Println("module: example.com/hello")
}
