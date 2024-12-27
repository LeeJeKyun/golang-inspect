package main

import "fmt"

func main() {
	x := [5]int{2, 35, 66, 35, 66}
	for i, v := range x {
		fmt.Println(i, "번째 값:", v)
	}
	fmt.Printf("%T\n", x)
}
