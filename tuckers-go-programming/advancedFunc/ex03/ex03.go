package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

// 반환타입이 func(int, int) int 인 함수(함수를 리턴하는 함수)
func getOperator(op string) func(int, int) int {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}
}

func main() {
	var operator func(int, int) int
	operator = getOperator("*")
	var result = operator(3, 4)
	fmt.Println(result)
}
