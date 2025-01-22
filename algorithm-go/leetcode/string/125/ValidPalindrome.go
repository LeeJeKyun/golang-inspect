package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	lower := strings.ToLower(s)
	chars := []rune(lower)

	lp := 0
	rp := len(lower) - 1

	for lp < rp {
		//문자나 넘버가 아닌동안 왼쪽 포인터 이동
		for lp < rp && !unicode.IsDigit(chars[lp]) && !unicode.IsLetter(chars[lp]) {
			lp++
		}
		//문자나 넘버가 아닌동안 우측 포인터 이동
		for lp < rp && !unicode.IsDigit(chars[rp]) && !unicode.IsLetter(chars[rp]) {
			rp--
		}
		//좌측, 우측 포인터가 동일하지 않으면 false리턴
		if chars[lp] != chars[rp] {
			fmt.Println("lp =", chars[lp], "rp =", chars[rp])
			return false
		}
		lp++
		rp--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("Do geese see God?"))
	r := 5560
	fmt.Printf("%T, %d, %c", r, r, r)
}
