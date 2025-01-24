package main

import (
	"fmt"
	"regexp"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	//단어-횟수 맵 생성
	m := make(map[string]int)

	//금지 단어집 생성 및 단어 추가
	set := make(map[string]bool)
	for _, v := range banned {
		set[v] = true
	}

	//소문자로 변경 및 특수문자 제거
	str := regexp.MustCompile(`\W+`).ReplaceAllString(paragraph, " ")
	lower := strings.ToLower(str)
	split := strings.Split(lower, " ")

	for _, v := range split {
		if v == "" || set[v] {
			continue
		}
		m[v]++
	}

	result := ""
	i := 0
	for k, v := range m {
		if v > i {
			result = k
			i = v
		}
	}

	return result
}

func main() {
	//paragraph := "Ross hit a ball, the hit BALL flew far away after it was hit."
	paragraph := "a."
	banned := []string{}
	fmt.Println(mostCommonWord(paragraph, banned))

	map1 := make(map[string]int)
	map1["key"] = 1
	fmt.Println("before", map1["key"])
	map1["key"] = 100
	fmt.Println("after", map1["key"])
	delete(map1, "key")
	fmt.Println("delete", map1["key"])
}
