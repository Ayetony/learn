package testing

import (
	"fmt"
	"testing"

	"example.com/hello/object"
)

func TestTypeInteger(t *testing.T) {
	var a1 object.Integer = 1
	var a2 object.Integer = 9
	fmt.Printf("a1.Less(a2): %v\n", a1.Less(a2))
	// 只有你需要修改对象的时候才会用到指针
	// go 语言的类型是基于值传递的，在修改变量的值并保留修改需要使用指针引用
	a1.Add(a2)
	fmt.Println("a = ", a1)
}
