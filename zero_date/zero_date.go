package main

import (
	"fmt"
	"time"
)

func main() {
	formatTimeStr := "0000-00-00 00:00:00"

	formatTime, err := time.Parse("2006-01-02 15:04:05", formatTimeStr)

	fmt.Println(formatTime, "aaa",err)
}
