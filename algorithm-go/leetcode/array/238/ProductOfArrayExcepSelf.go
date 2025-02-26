package main

import "fmt"

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	p := 1
	for i := 0; i < len(nums); i++ {
		result[i] = p
		p *= nums[i]
	}
	p = 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= p
		p *= nums[i]
	}

	return result
}

func main() {
	input := []int{1, 3, 5, 7}
	result := productExceptSelf(input)
	output := []int{105, 35, 21, 15}
	for i, v := range output {
		if v == result[i] {
			fmt.Println("Success")
		} else {
			fmt.Println("Fail!")
		}
	}

}
