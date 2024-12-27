package main

import (
	"fmt"
	"strconv"
)

func main() {
	var first string
	fmt.Scan(&first)
	num, _ := strconv.Atoi(first)

	var maxNum float32 = 0
	scoreArr := make([]float32, num)
	var totalScore float32 = 0
	for i := 0; i < len(scoreArr); i++ {
		fmt.Scan(&scoreArr[i])
		if scoreArr[i] > maxNum {
			maxNum = scoreArr[i]
		}
		totalScore += scoreArr[i]
	}
	fmt.Println(totalScore / float32(num) / maxNum * float32(100))

}
