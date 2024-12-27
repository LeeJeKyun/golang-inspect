package main

import "fmt"

// opFunc를 func(a, b int) int 함수정의의 별칭으로 선언
type opFunc func(a, b int) int

func getOperator(op string) opFunc {
	if op == "+" {
		return func(a, b int) int {
			return a + b
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b
		}
	} else {
		return nil
	}
}

func main() {
	fn := getOperator("*")

	result := fn(3, 4)
	fmt.Println(result)
}
