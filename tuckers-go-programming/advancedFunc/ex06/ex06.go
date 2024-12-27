package main

import "fmt"

func CaptureLoop() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop2")
	for i := 0; i < 3; i++ {
		v := i
		f[i] = func() {
			fmt.Println(v)
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

// Golang 1.22버전부터 클로저 내부에 변수할당이 이뤄지게끔 변경되면서 CaptureLoop와 CaptureLoop2가 동일한 동작을 보인다.
func main() {
	CaptureLoop()
	CaptureLoop2()
}
