package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "Hello, playground")
	io.WriteString(os.Stdout, "Hello, playground")
	n, err := fmt.Fscanln(os.Stdin)
	if err != nil {
		fmt.Println(string(n))
	}
}
