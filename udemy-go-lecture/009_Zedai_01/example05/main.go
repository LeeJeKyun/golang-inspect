package main

import "fmt"

type hotdog int

var x hotdog
var y int

func main() {
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
