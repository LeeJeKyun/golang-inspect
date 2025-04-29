package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("sample.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString("Hello, File!")

	read, err := os.Open("sample.txt")
	if err != nil {
		panic(err)
	}
	defer read.Close()

	content, err := io.ReadAll(read)
	if err != nil {
		panic(err)
	}
	fmt.Println("파일 내용:", string(content))
}
