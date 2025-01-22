package main

import "fmt"

func reverseString(s []byte) {

	start := 0
	end := len(s) - 1

	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

func main() {
	bytes := []byte{'r', 'a', 'c', 'e', 'c', 'a', 'r'}
	reverseString(bytes)
	s := string(bytes)
	fmt.Println(s)
}
