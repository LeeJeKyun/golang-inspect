package main

import "fmt"

// defer가 포함된 함수가 종료될 때마다 실행된다.
// 여러개의 defer 함수가 있을 경우 맨아래에서부터 위로 올라오며 순서대로 실행된다.
func main() {
	defer bar()
	defer foo()
	foo()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
