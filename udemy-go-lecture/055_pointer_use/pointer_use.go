package main

import "fmt"

func main() {
	x := 0
	noPointer(x)
	fmt.Println(x)

	fmt.Println()

	y := 0
	pointer(&y)
	fmt.Println(y)
}

func noPointer(y int) {
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}

func pointer(y *int) {
	fmt.Println(*y)
	*y = 44
	fmt.Println(*y)
}
