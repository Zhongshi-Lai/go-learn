package main

import "fmt"

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
	for i := 0; i < n; i++ {
		fmt.Println(chineseStr[i])
	}

}
