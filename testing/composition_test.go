package testing

import (
	"fmt"
	"testing"

	"example.com/hello/object"
)

func TestComposition(t *testing.T) {
	// 初始化
	var y = &object.Y{
		Logger:  &object.Logger{Level: 3},
		Command: "anonymous assign",
	}
	fmt.Println(y.Command)
	//如果没有初始化，则出错
	y.Level = 4
	y.Log()
}
