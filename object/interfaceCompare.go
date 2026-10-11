package object

import "fmt"

type ReadWriter interface {
	Read(buf []byte) (n int, err error)
	Writer(buf []byte) (n int, err error)
}

type IStream interface {
	Read(buf []byte) (n int, err error)
	Writer(buf []byte) (n int, err error)
}

type Writer interface {
	Read(buf []byte) (n int, err error)
}

// CompareInterface 接口赋值
// 将一个接口赋值给另一个接口，只要两个接口拥有相同的方法列表（次序不要紧）
// 如果一个接口实现的方法是另外一个接口实现方法的子集，后者可以传递给前者，相反则不可以
func CompareInterface() {
	var file1 = new(ReadWriter)
	var file2 IStream = *file1
	var file3 Writer = file2
	fmt.Print(file2, file3)
}
