package main

import "fmt"

func Outer() func() {
	fmt.Println("Outer Function Call!")
	return func() {
		fmt.Println("Return Function Call")
	}
}

func main() {
	f := Outer()
	f()
}
