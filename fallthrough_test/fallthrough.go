package main

import "fmt"

func main() {

	i := 1
	switch {
	case i == 1:
		fmt.Println("into case 1")
		fallthrough
	case i == 2:
		// 所以fallthrough 会直接进入下一个case,并不会再次判断下一个case的条件了
		fmt.Println("into case 2")
		fmt.Println(i)
	case i == 3:
		fmt.Println("ssss")
	}
}
