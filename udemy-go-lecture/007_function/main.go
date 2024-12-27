package main

import "fmt"

func main() {
	fmt.Println("abc")
	boolean := NameOfFunc(2, 2)
	fmt.Println(boolean)
}

func NameOfFunc(a int, b int) (isTrue bool) {
	isTrue = (a == b)
	return
}
