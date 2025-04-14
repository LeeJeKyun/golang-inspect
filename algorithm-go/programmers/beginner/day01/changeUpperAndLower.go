package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s1 string
	fmt.Scan(&s1)
	rs := []rune{}
	for _, v := range s1 {
		var r rune
		if unicode.IsUpper(v) {
			r = unicode.ToLower(v)
		} else {
			r = unicode.ToUpper(v)
		}
		rs = append(rs, r)
	}
	fmt.Println(string(rs))
}
