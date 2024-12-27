package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Println(x[1:3])

	for i, v := range x {
		fmt.Println(i, v)
	}

	for j := 0; j <= 4; j++ {
		fmt.Println(j, x[j])
	}
}
