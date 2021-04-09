package main

import "fmt"

func main() {

JLOOP:
	for j := 0; j < 5; j++ {
	// ILOOP:
		for i := 0; i < 10; i++ {
			if i > 5 {
				break JLOOP
			}
			fmt.Println(i)
		}

	}
}
