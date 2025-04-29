package main

import "fmt"

// 클로저(Closure)는 함수 리터럴과 그 함수가 참조하는 외부 변수의 환경을 함께 캡슐화한 것
// 자신이 선언된 외부 범위의 변수에 접근할 수 있는 함수
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	f := adder()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}
