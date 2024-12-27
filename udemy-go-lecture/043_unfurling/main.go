package main

import "fmt"

func main() {
	x := sum("james", 4, 5, 6)
	fmt.Println("The Total is", x)
}
func sum(s string, x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position,", i, "we are now adding,", v, "to the total which is now,", sum)
	}
	fmt.Println("The total is,", sum)
	return sum
}
