package main

import (
	"fmt"
	"strconv"
)

func main() {
	var A, B, C int
	fmt.Scan(&A, &B, &C)

	fmt.Println(A + B - C)

	a := strconv.Itoa(A)
	b := strconv.Itoa(B)
	str := a + b
	atoi, _ := strconv.Atoi(str)
	fmt.Println(atoi - C)
}
