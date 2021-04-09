package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(100000 + rand.Intn(999999-100000))
}
