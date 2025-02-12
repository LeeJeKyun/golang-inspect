package main

import "fmt"

func trap(height []int) int {
	//결과값(볼륨)
	volume := 0
	//좌포인터
	lp := 0
	//우포인터
	rp := len(height) - 1
	//좌최대값
	lM := height[0]
	//우최대값
	rM := height[rp]
	for lp < rp {
		if height[lp] > lM {
			lM = height[lp]
		}
		if height[rp] > rM {
			rM = height[rp]
		}
		if lM < rM {
			volume += lM - height[lp]
			lp++
		} else {
			volume += rM - height[rp]
			rp--
		}
	}

	return volume
}

func main() {
	input := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	i := trap(input)
	fmt.Println(i)
}
