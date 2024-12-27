package main

import "fmt"

func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
func foo2(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := foo(ii...)
	fmt.Println("foo =", n)
	fmt.Println("foo2 =", foo2(ii))
}
