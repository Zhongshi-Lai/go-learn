package main

import (
	"fmt"
	"math"
)

func isEqual(f1, f2, p float64) bool {
	return math.Abs(f1-f2) < p
}

func main() {

	// 直接比较可能导致不稳定的结果
	a := 10.0
	b := 10.0

	if a == b {
		fmt.Println("===float===")
		fmt.Print(a)
		fmt.Print(b)
		fmt.Println("===float===")
	}

	c := 10
	d := 10

	if c == d {
		fmt.Println("===int===")
		fmt.Print(c)
		fmt.Print(d)
		fmt.Println("===int===")
	}

	// p 自定义一个精度
	p := 0.001
	res := isEqual(a, b, p)
	fmt.Println(res)

}
