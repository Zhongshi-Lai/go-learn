package main

import (
	"fmt"
	"strings"

	uuid "github.com/satori/go.uuid"
)

func main() {

	b := strings.Replace(uuid.NewV4().String(), "-", "", -1)
	fmt.Println(b)
}
