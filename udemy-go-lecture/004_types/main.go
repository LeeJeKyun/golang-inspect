package main

import "fmt"

var y = 42

// 'z'라는 변수 선언 및 "Shaken, not stirred"라는 문자열 타입의 값 대입
// 문자열 타입이다.

var z string

var a string = `james said,
" Shaken,

not stirred"`

// JS와 같은 동적 프로그래밍 언어는 아래와 같은 대입이 가능했을것.
// Golang은 정적 프로그래밍 언어이다.
// 정적 프로그래밍 언어에서 변수는 정해진 타입의 값만 대입이 가능하다.

func main() {
	z = "str"
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	// 문자열 타입이기 때문에 정수 타입 대입 불가
	// z = 43
	// fmt.Println(z)
	// fmt.Printf("%T\n", z)

}
