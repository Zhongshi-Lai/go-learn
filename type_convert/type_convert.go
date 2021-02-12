package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	// 断言
	// s, ok := x.(T)
	// 把 x 转换为 T类型 s是接收的新类型的变量
	_, err := os.Open("test_file")

	if err != nil {
		if e, ok := err.(*os.PathError); ok && e.Err != nil {
			fmt.Println("=== type assert success ===")
			fmt.Println(e)
			fmt.Println(reflect.TypeOf(e))
			fmt.Println("=== type assert success ===")
		}
	}

	// switch 转换类型
	switch i := err.(type) {
	// type of i is type of x (interface{})
	default:
		fmt.Println(i)
	}

}
