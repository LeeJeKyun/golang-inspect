package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//bufferd Bytes
	str := "Hello, Writer!"
	bufferd := []byte(str)

	//Writer(File) 생성
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//Write
	file.Write(bufferd)

	//Write된 파일 확인
	open, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer open.Close()

	all, err := io.ReadAll(open)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(all))
}
