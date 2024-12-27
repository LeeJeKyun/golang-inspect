package main

import (
	"fmt"
	"sync"
)

// 아래 선언으로 제로값으로 초기화된 WaitGroup 인스턴스가 생성된다고 볼 수 있다.
// 고랭의 특징으로 변수선언과 동시에 제로값으로 초기화됨.
// 따라서 아래 main문과 SumAtoB함수에서 wg의 메서드를 사용할 수 있음.
var wg sync.WaitGroup

// 정수 a,b를 매개변수로 받아서 a~b사이에 있는 정수를 모두 더한 후 출력하는 함수
func SumAtoB(a, b, num int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d번째 고루틴%d부터 %d까지 합계는 %d입니다.\n", num+1, a, b, sum)
	wg.Done()
}

func main() {

	wg.Add(10)
	//1~1000000000까지의 수를 모두 더하는 함수를 동시적으로 10번 반복함.
	for i := 0; i < 10; i++ {
		go SumAtoB(1, 1000000000, i)
	}

	wg.Wait()
	fmt.Println("모든 계산이 완료되었습니다.")
}
