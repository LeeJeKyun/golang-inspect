package main

import "fmt"

func main() {
	slice1 := make([]string, 5)

	slice2 := make([]int, 3, 5)

	slice3 := []string{"a", "b", "c", "d", "e"}

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	var sliceNil []int
	fmt.Println(sliceNil, len(sliceNil), cap(sliceNil))

	newSlice := append(slice3, "30")
	fmt.Println(newSlice)

}
