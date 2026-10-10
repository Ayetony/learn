package testing

import (
	"fmt"
	"testing"

	"example.com/hello/object"
)

func TestArrayRef(t *testing.T) {
	arr := []int{1, 2, 2}
	// 不会改变arr的值，而是复制了一份
	object.ModifyArray([3]int(arr))
	fmt.Println(arr)
}

func TestRefArrayWithSlice(t *testing.T) {
	slices := []int{1, 2, 3}
	object.ModifySliceArray(slices)
	fmt.Println(slices)
}
