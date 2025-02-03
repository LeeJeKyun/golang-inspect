package main

import "fmt"

var start, maxLen int

func longestPalindrome(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}

	start, maxLen = 0, 0

	for i := 0; i < length-1; i++ {
		extendPalindrome(s, i, i)
		extendPalindrome(s, i, i+1)
	}
	return s[start : start+maxLen]
}

func extendPalindrome(s string, l, r int) { //문자열, 좌, 우 시작포인터 전달
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
		if maxLen < r-l-1 {
			start = l + 1
			maxLen = r - l - 1
		}
	}
}

func main() {
	s := "cbbd"
	palindrome := longestPalindrome(s)
	fmt.Println(palindrome)
}
