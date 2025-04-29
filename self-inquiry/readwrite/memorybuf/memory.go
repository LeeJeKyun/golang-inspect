package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var buf bytes.Buffer

	buf.WriteString("Hello, MemoryBuffer!")

	content, err := io.ReadAll(&buf)
	if err != nil {
		panic(err)
	}

	fmt.Println("메모리 내용:", string(content))
}
