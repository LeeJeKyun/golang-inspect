package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int) {
	for n := range ch {
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	//채널을 제때 닫아주지 않아서 생기는 데드락 = 고루틴leak
	//close(ch)
	wg.Wait()
}
