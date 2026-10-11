package transport

import "fmt"

func add(ch chan int) {
	ch <- 1
}

func ExecChanAdd(times int) {

	chs := make([]chan int, times)

	for i := 0; i < times; i++ {
		chs[i] = make(chan int)
		go add(chs[i])
	}

	for _, ch := range chs {
		value, exist := <-ch
		fmt.Println(value, exist)
	}

}
