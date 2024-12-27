package main

import "fmt"

func main() {
	fmt.Println(4 * 3 * 2 * 1)
	n := factorial(4)
	fmt.Println("Factoriial:", n)
	n2 := factorialLoop(4)
	fmt.Println("FactorialLoop:", n2)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func factorialLoop(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
