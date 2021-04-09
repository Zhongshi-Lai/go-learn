package main

import (
	"fmt"
	"reflect"
)

type TestA struct {
	Flag bool
}

func (t *TestA) Run(a ...int) int {
	var out int = 0
	for _, i := range a {
		out += i
	}
	return out
}

/// 注意: reflect.ValueOf接收的是值
func testA() {
	i := TestA{}
	vl := reflect.ValueOf(i)
	fmt.Println("value:", vl.FieldByName("Flag").Interface())
	//
	vl4fun := reflect.ValueOf(&i)
	args := []reflect.Value{reflect.ValueOf(int(2))}
	fmt.Println("value:", vl4fun.MethodByName("Run").Call(args)[0].Interface())
}

/// 注意: reflect.ValueOf接收的是指针
func testA4P() {
	i := new(TestA)
	vl := reflect.ValueOf(i)
	fmt.Println("value:", vl.Elem().FieldByName("Flag").Interface())
	//
	vl4fun := reflect.ValueOf(i)
	args := []reflect.Value{reflect.ValueOf(int(2))}
	fmt.Println("value:", vl4fun.MethodByName("Run").Call(args)[0].Interface())
}

///TestB
type TestB struct {
	Flag bool
}

/**************************************************************************/
// 注意: 此处定义是非指针
func (t TestB) Run(a ...int) int {
	var out int = 0
	for _, i := range a {
		out += i
	}
	return out
}
func testB() {
	i := TestB{}
	vl := reflect.ValueOf(i)
	//
	fmt.Println("value:", vl.FieldByName("Flag").Interface())
	args := []reflect.Value{reflect.ValueOf(int(2))}
	fmt.Println("value:", vl.MethodByName("Run").Call(args)[0].Interface())
}

//main
func main() {
	testB()
}
