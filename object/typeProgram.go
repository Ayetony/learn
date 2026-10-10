// Package object 类型系统编程
package object

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}
