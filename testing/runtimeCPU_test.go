package testing

import (
	"fmt"
	"runtime"
	"testing"
)

func TestNumOfCPU(t *testing.T) {
	fmt.Println(runtime.NumCPU())
}
