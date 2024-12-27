package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	var input int
	money := 1000
	for {
		if money <= 0 || money >= 5000 {
			fmt.Println("종료되었습니다.")
			break
		}
		n := rand.IntN(5)
		n += 1
		fmt.Println(n)
		fmt.Printf("1 ~ 5 사이 값을 입력하세요: ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			stdin.ReadString('\n')
		}

		if n == input {
			money += 500
			fmt.Println("축하합니다. 현재 잔액:", money)
		} else if n != input {
			money -= 100
			fmt.Println("아쉽습니다.. 현재 잔액:", money)
		}

	}

}
