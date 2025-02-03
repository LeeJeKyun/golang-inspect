package main

import "fmt"

var left, maxLen int

func longestPalindrome(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}

	for i := 0; i < length-1; i++ {
		fmt.Println(left, maxLen)
		extendPalindrome(s, i, i+1)
		extendPalindrome(s, i, i+2)
	}
	return s[left : left+maxLen]
}

func extendPalindrome(s string, l, r int) { //문자열, 좌, 우 시작포인터 전달
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
		if maxLen < r-l-1 {
			left = l + 1
			maxLen = r - l - 1
		}
	}
}

func main() {
	s := "cbbd"
	palindrome := longestPalindrome(s)
	fmt.Println(palindrome)
}
