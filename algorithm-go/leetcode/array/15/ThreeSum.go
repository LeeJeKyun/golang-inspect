package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	// 길이만큼 반복
	for i := 0; i < len(nums); i++ {
		//좌, 우 포인터 선언
		lp := i + 1
		rp := len(nums) - 1
		//만약 이전 인덱스의 값과 동일할 경우 continue
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		//좌, 우 포인터가 유지될 동안 반복
		for lp < rp {
			//만약 l, i, r 세 값의 합이 0이면 인덱스 추출하여 result에 삽입 후 continue
			sum := nums[lp] + nums[i] + nums[rp]
			if sum == 0 {
				result = append(result, []int{nums[lp], nums[i], nums[rp]})
				//lp < rp이고 값이 동일할 경우 lp++
				for lp < rp && nums[lp] == nums[lp+1] {
					lp++
				}
				for lp < rp && nums[rp] == nums[rp-1] {
					rp--
				}
				lp++
				rp--
				//세 값의 합이 0보다 크면 r--
			} else if sum > 0 {
				rp--
				//세 값의 합이 0보다 작으면 l++
			} else {
				lp++
			}
		}
	}

	return result
}

func main() {
	input := []int{-1, 0, 1, 2, -1, -4}
	sum := threeSum(input)
	fmt.Println(sum)
}
