package main

import (
	"fmt"
	"reflect"
)

type Data struct {
	a, b int
}

func main() {
	var num int = 1
	fmt.Println(reflect.TypeOf(num))

	var s string = "Hello, world!"
	fmt.Println(reflect.TypeOf(s))

	var f float32 = 1.3
	fmt.Println(reflect.TypeOf(f))

	var data Data = Data{1, 2}
	fmt.Println(reflect.TypeOf(data))
}
