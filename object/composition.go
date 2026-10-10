package object

import (
	"fmt"
	"log"
)

type Job struct {
	Command string
	*log.Logger
}

type Logger struct {
	Level int
}

type Y struct {
	*Logger
	Command string
}

// Start job 可以直接使用 Logger 下的方法
// anonymous composition
func (job *Job) Start() {
	job.Print("hello Job")
}

func (y *Y) Log() {
	fmt.Println(y.Level)
}
