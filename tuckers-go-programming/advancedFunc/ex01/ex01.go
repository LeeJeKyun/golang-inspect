package main

import "fmt"

func sum(nums ...int) int {
	sum := 0

	fmt.Printf("nums 타입: %T\n", nums)
	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(10, 20))
	fmt.Println(sum())

	xi := []int{1, 3, 5, 7, 9}
	//unfurling
	fmt.Println(sum(xi...))
}
