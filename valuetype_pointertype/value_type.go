package main

import "fmt"

func testStringFuncParams(s string) {
	fmt.Println("===in func ===")
	fmt.Println(&s)
	fmt.Println("===in func ===")
}

func testPointerFuncParams(s *string) {
	fmt.Println("===in func pointer ===")
	fmt.Println(s)
	fmt.Println(*s)
	// fmt.Println(*s + "1")
	*s += "1"
	fmt.Println(s)
	fmt.Println("===in func pointer ===")
}

func main() {
	// 值类型
	// 本质是内存地址的变动,传入func之后是另一个内存地址了
	// int byte bool float string
	// array struct pointer

	a := "hello"
	b := &a


	fmt.Println("===before func ===")
	fmt.Println(&a)
	fmt.Println("===before func ===")
	testStringFuncParams(a)

	// TODO(laizhongshi):这块我迷茫了,指针到底是值语义还是引用语义
	testPointerFuncParams(&a)

	fmt.Println(b)
	fmt.Println(a)
	// 引用类型 slice map channel interface
}
