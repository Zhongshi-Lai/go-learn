package main

import (
	"fmt"
	"time"
)

func deferLog() {
	r := recover()
	if r == nil {
		return
	}
	fmt.Println("func panic")
}

func goroutinePanic() {
	defer deferLog()
	panic("aaa")

}

func main() {

	// defer deferLog()

	go func() {
		defer deferLog()
		goroutinePanic()
	}()

	time.Sleep(time.Duration(200) * time.Millisecond)

}
