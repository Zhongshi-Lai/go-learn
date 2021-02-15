package main

import "fmt"

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
	// 注意要使用&a 因为add方法使用了a *Integer
	// 如果使用 var b LessAdder = a 则不会自动生成Add成员方法
	// 但是如果使用var b LessAdder = &a 会自动生成一个Less方法,使用指针
	var b LessAdder = &a

	fmt.Println(b)

	// Lesser接口就可以
	var c Lesser = &a
	var d Lesser = a

	fmt.Println(c, d)

	// 接口赋值只能超集向子集转换

	// 对象赋值给接口
	// 接口赋值给接口

	var f Lesser = b
	fmt.Println(f)

}
