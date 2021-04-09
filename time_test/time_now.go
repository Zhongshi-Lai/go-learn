package main

import (
	"fmt"
	"time"
)

// 1615862265.4143848
// 1615862370.378985000

// 1615862265414

func main() {
	now := time.Now().UnixNano() / 1e6

	fmt.Println(now)
}
