package main

import "fmt"

func twoSum(nums []int, target int) []int {
	//결과값 생성
	r := make([]int, 0)
	//target-값:index 맵 생성
	m := make(map[int]int)
	for i, v := range nums {
		m[target-v] = i
	}
	// 답이 중복되는 경우 제약조건
	// 값이 아닌 인덱스 저장

	for i, v := range nums {
		j, ok := m[v]
		if ok {
			if i != j {
				r = append(r, j)
				r = append(r, i)
				return r
			}
		}
	}
	return nil
}

func main() {
	input := []int{2, 6, 10, 11}
	target := 8
	sum := twoSum(input, target)
	fmt.Println(sum)

	m1 := make(map[int][]int)
	v, ok := m1[33]
	fmt.Println(v == nil)
	fmt.Println(ok)
}
