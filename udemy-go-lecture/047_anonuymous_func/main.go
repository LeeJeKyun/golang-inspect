package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("Anonymous func ran")
	}()
	func(x int) {
		fmt.Println("The meaning of life:", x)
	}(43)
}

func foo() {
	fmt.Println("foo ran")
}
