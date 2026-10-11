package transport

import (
	"fmt"
	"sync"
)

var Counter = 0

func Count(lock *sync.Mutex) {

	lock.Lock()
	Counter++
	lock.Unlock()
	fmt.Println("Counter value is ", Counter)

}
