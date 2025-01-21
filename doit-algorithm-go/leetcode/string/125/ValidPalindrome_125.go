package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	fields := strings.Index("ABC", s)
	fmt.Println(fields)
	return true
}

func main() {
	isPalindrome("ABCDE")
}
