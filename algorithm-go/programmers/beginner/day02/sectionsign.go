package main

import (
	"fmt"
	"slices"
)

func solution(emergency []int) []int {
	increase := 1
	for i := 0; i < len(emergency); i++ {
		i2 := slices.Max(emergency)
		emergency[i2] = increase
		increase++
	}

	return emergency
}
func main() {
	emergency := []int{3, 76, 24}
	i := solution(emergency)
	fmt.Println(i)
}
