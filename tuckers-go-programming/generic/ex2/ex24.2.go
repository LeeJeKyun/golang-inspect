package main

import "fmt"

func Print[T1 any, T2 any](a T1, b T2) {
	fmt.Println(a, b)
}

func Print2(a, b interface{}) {
	fmt.Println(a, b)
}

func main() {
	Print(1, 2)
	Print(3.14, 1.43)
	Print("Hello", "World")
	Print(1, "Hello")

	Print2(1, 2)
	Print2(3.14, 1.43)
	Print2("Hello", "World")
	Print2(1, "Hello")
}
