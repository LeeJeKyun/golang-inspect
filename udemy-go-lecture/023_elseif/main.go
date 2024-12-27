package main

import "fmt"

func main() {
	x := 42
	if x == 40 {
		fmt.Println("out value was 40")
	} else if y := 42; y == 42 && x == 41 {
		fmt.Println("out value was 41")
	} else {
		fmt.Println("out value was not 40")
	}
}
