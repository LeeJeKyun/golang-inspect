package main

import "sort"

func arrayPairSum(nums []int) int {
	result := 0
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < len(nums); i = i + 2 {
		result += nums[i]
	}
	return result
}
