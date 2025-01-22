package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func reorderLogFiles(logs []string) []string {
	digits := make([]string, 0)
	letters := make([]string, 0)

	for _, v := range logs {
		split := strings.Split(v, " ")
		runes := []rune(split[1])
		if unicode.IsDigit(runes[0]) {
			digits = append(digits, v)
		} else {
			letters = append(letters, v)
		}
	}

	sort.Slice(letters, func(i, j int) bool {
		s1 := letters[i]
		s2 := letters[j]

		s1x := strings.Split(s1, " ")
		s2x := strings.Split(s2, " ")
		f1 := strings.Join(s1x[1:], " ")
		f2 := strings.Join(s2x[1:], " ")
		fmt.Println(f1)
		fmt.Println(f2)
		if f1 == f2 {
			return s1x[0] < s2x[0]
		} else {
			return f1 < f2
		}
	})

	letters = append(letters, digits...)

	return letters
}

func main() {
	input := []string{"id1 8 1 5 1", "id2 art can", "id3 3 6", "id4 own kit dig", "id5 art zero"}
	fmt.Println(reorderLogFiles(input))
}
