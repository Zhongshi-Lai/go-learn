package main

import "fmt"

type Student struct {
	name string
}

func (s Student) tellName() {
	fmt.Println(s.name)
}

func (s Student) alterName() {
	s.name = "xiaoming gai"
	fmt.Println("aaa")
}

func (s *Student) alterNameV2() {
	s.name = "xiaoming gai"
}

func main() {

	s := Student{
		name: "xiaoming",
	}
	s.tellName()

	s.alterName()

	s.tellName()

	s.alterNameV2()
	s.tellName()
}
