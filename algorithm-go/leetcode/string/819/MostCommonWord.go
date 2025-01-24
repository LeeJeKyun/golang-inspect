package main

import (
	"fmt"
	"strings"
	"unicode"
)

func mostCommonWord(paragraph string, banned []string) string {
	//m := make(map[string]int)
	//ban := make(map[string]int)

	//특수 문자를 포함하여 모두 공백문자로 변경
	runes := []rune(paragraph)
	for i, r := range runes {
		if !unicode.IsLetter(r) {
			runes[i] = ' '
		} else {
			runes[i] = r
		}
	}
	s := string(runes)

	return ""
}

func main() {
	paragraph := "Ross hit a ball, the hit BALL flew far away after it was hit."
	banned := []string{"hit"}
	fmt.Println(mostCommonWord(paragraph, banned))
}
