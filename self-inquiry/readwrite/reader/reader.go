package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// create dataSource
	dataSource := "Hello, Reader!"

	// Create Reader
	src := strings.NewReader(dataSource)

	// create a packet
	p := make([]byte, 3)

	// read 'dataSource' until an error is returned
	for {
		// read 'p'bytes from 'src'
		n, err := src.Read(p)
		fmt.Printf("%d bytes read, data: %s\n", n, p[:n])

		if err == io.EOF {
			fmt.Println("--end-of-file--")
			break
		} else if err != nil {
			fmt.Println("Some error occured!", err)
			break
		}
	}
}
