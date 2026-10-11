package object

import (
	"fmt"
	"reflect"
)

func ModifyArray(arr [3]int) {
	// 指向具体的位置，就是值传递，不影响外部
	arr[1] = 2
}

func ModifySliceArray(s []int) {
	// 未指定对应的索引则是一个指针引用
	fmt.Println("s type is ", reflect.ValueOf(s).Type(), reflect.TypeOf(s).String())
	s[1] = 10
}
