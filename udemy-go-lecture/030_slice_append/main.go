package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	for i := 0; i < len(y); i++ {
		x = append(x, y[i])
	}
	fmt.Println(x)
}
