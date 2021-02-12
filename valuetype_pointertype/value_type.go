package main

import "fmt"

func testStringFuncParams(s string) {
	fmt.Println("===in func ===")
	fmt.Println(&s)
	fmt.Println("===in func ===")
}

func testPointerFuncParams(s *string) {
	fmt.Println("===in func pointer ===")
	fmt.Println(*s)
	fmt.Println(*s + "1")
	fmt.Println("===in func pointer ===")
}

func main() {
	// 值类型
	// 本质是内存地址的变动,传入func之后是另一个内存地址了
	// int byte bool float string
	// array struct pointer
	a := "hello"
	fmt.Println("===before func ===")
	fmt.Println(&a)
	fmt.Println("===before func ===")
	testStringFuncParams(a)

	// 即使传输指针进去,修改*的值,也不会影响外部的
	testPointerFuncParams(&a)

	b := &a
	fmt.Println(b)
	fmt.Println(a)
	// 引用类型 slice map channel interface
}
