package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	// 맵 생성
	m := make(map[string][]string)
	result := make([][]string, 0)
	// 배열 요소마다 반복
	for _, v := range strs {
		//문자열 정렬
		runes := []rune(v)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		sorted := string(runes)
		_, ok := m[sorted]
		if !ok {
			m[sorted] = make([]string, 0)
		}
		m[sorted] = append(m[sorted], v)
	}

	for _, v := range m {
		result = append(result, v)
	}

	return result
}

func main() {
	input := []string{"eat", "tea", "tan", "ate", "ant", "cat"}
	fmt.Println(groupAnagrams(input))
}
