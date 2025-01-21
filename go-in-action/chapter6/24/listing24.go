package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4  //실행할 고루틴의 개수
	taskLoad         = 10 //처리할 작업의 개수
)

// 종료시까지 대기할 WaitGroup
var wg sync.WaitGroup

func init() {
	// 랜덤값 생성기 초기화
	rand.NewSource(time.Now().Unix())
}

func main() {
	// 작업 부하를 관리하기 위한 버퍼가 있는 채널을 생성
	tasks := make(chan string, taskLoad)

	// 작업을 처리할 고루틴을 실행
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// 실행할 작업을 추가한다.
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("작업: %d", post)
	}

	// 작업을 모두 처리하면
	// 채널을 닫는다.
	close(tasks)

	// 모든 작업이 처리될 때까지 대기한다.
	wg.Wait()
}

func worker(tasks chan string, worker int) {
	// 함수가 리턴될 때 Done함수 호출 예약
	defer wg.Done()

	//
	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("작업자: %d : 종료합니다.\n", worker)
			return
		}

		fmt.Printf("작업자: %d : 작업 시작: %s\n", worker, task)

		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("작업자: %d : 작업 완료: %s\n", worker, task)
	}
}
