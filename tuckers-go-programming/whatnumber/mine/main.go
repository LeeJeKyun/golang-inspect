package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := 0
	i := rand.Intn(100)
	var input int

	for {
		n++
		fmt.Printf("숫자값을 입력하세요:")
		fmt.Scanln(&input)

		if input > i {
			fmt.Println("입력하신 숫자가 더 큽니다.")
		} else if input < i {
			fmt.Println("입력하신 숫자가 더 작습니다.")
		} else {
			fmt.Println("숫자를 맞췄습니다. 축하합니다. 시도한 횟수:", n)
			break
		}
	}
}
