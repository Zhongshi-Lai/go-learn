package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var chineseStr = "hello 世界"

	// range 法遍历
	for _, each := range chineseStr {
		// fmt.Println(index)
		fmt.Println(each)
	}

	fmt.Println("=========")

	// 下标法遍历,中文字符占3个
	n := len(chineseStr)
	fmt.Println("len<<<", n)
	nStr := utf8.RuneCountInString(chineseStr)
	fmt.Println("lenstr<<<", nStr)
	for i := 0; i < n; i++ {
		fmt.Println(chineseStr[i])
	}

}
