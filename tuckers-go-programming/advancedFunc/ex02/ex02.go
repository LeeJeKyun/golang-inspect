package main

import (
	"fmt"
	"os"
)

// defer함수는 선언된 함수가 종료될때 역순으로 실행된다.
func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}

	defer fmt.Println("반드시 호출됩니다") //3번
	defer f.Close()                //2번
	defer fmt.Println("파일을 닫았습니다") //1번

	fmt.Println("파일에 Hello World를 씁니다.")
	fmt.Fprintln(f, "Hello Golang World")
}
