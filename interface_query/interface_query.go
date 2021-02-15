package main

import (
	"fmt"
	"reflect"
)

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
}

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

type Lesser interface {
	Less(b Integer) bool
}

func main() {
	var a Integer = 1

	var b LessAdder = &a

	// TODO(laizhongshi):本质是类型转换吧?
	// 转换后的c变量,就只有Less方法了
	if c, ok := b.(Lesser); ok {
		fmt.Println(ok)
		fmt.Println(c)
		d := reflect.TypeOf(c)
		fmt.Println(d)
	}

}
