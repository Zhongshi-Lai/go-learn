package main

import "fmt"

// float 强行转int

func main() {
	a := 1191
	b := int(a / 100)
	fmt.Println(b)
}
