package main

import "fmt"

func main() {
	f := func(x, y int) {
		fmt.Println(x * y)
	}

	var g func()

	g = func() {
		fmt.Println("g() function call")
	}

	f(3, 7)
	fmt.Printf("%T\n", f)
	fmt.Println("done")
	g()
}
