package main

import "fmt"

// 'y'라는 변수를 선언하고 43이라는 값을 대입한다.
// 선언&대입은 초기화와 같다.
var y = 43

// z라는 변수 선언 및 z의 타입은 int라고 선언
// boolean에는 false, 숫자값(int, float)에는 0, 문자열에는 "", 포인터, 함수, 인터페이스, 슬라이스, 채널, 맵에는 nil이 기본값으로 들어간다.
// 출처 : https://go.dev/ref/spec#The_zero_value
var z int

func main() {
	//단축 선언 연산자 - 함수내에서만 사용가능(전역변수 사용 x)
	//변수 선언 & 특정 타입의 값 대입
	x := 42
	fmt.Println(x)

	fmt.Println(y)
	foo()

	fmt.Println(z)
}

func foo() {
	fmt.Println("agein:", y)
}
