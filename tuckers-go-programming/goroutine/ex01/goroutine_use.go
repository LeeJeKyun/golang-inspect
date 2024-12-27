package main

import (
	"fmt"
	"time"
)

// 300밀리초마다 한글을 프린트하는 함수
func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}
}

// 400밀리초마다 숫자를 프린트하는 함수
func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func main() {
	//고루틴을 이용하여 각기 다른 고루틴에서 한글과 숫자가 프린트됨.
	go PrintHangul()
	go PrintNumbers()

	//메인 고루틴이 종료될 경우 자식 고루틴도 모두 종료되므로 해당 부분을 방지하기 위해 슬립을 걸어줌.
	//이후 코드에선 sync.WaitGroup을 이용해 하위 고루틴이 모두 종료될때까지 기다려준다.
	time.Sleep(3 * time.Second)
}
