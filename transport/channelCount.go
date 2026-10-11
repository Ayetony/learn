package transport

import "fmt"

func add(ch chan int) {
	ch <- 1
}

func ExecChanAdd(times int) {

	chs := make([]chan int, times)

	// range 只是返回副本，如果使用副本，下次再使用就没有值，则 <-ch 永远接收不到，因为 ch 变量并没有改变
	for i := range chs {
		chs[i] = make(chan int)
		go add(chs[i])
	}

	for _, ch := range chs {
		value, exist := <-ch
		fmt.Println(value, exist)
	}

}
