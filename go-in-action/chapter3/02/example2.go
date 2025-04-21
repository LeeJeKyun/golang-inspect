package main

import (
	"fmt"
	"go-in-action-self/go-in-action/chapter3/02/inittest"
	"strings"
)

func main() {
	a := "a"
	b := "b"
	//compare 함수는 내림차순 -1, 동일 0, 오름차순 1을 반환한다.
	compare := strings.Compare(b, a)
	fmt.Println("compare =", compare)

	s := "Hello World"
	sub := "Hello`"

	contains := strings.Contains(s, sub)
	fmt.Println("contains =", contains)

	chars := []byte(s)
	for _, c := range chars {
		char := fmt.Sprintf("%c", c)
		fmt.Println(char)
	}

	inittest.PrintHello()

}
