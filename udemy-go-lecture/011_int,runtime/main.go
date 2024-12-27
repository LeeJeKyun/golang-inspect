package main

import (
	"fmt"
	"runtime"
)

// var x int
// var y float64

func main() {

	// x = 2147483648
	// y = 42.34534
	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Printf("%T\n", x)
	// fmt.Printf("%T", y)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
