package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func printer(a int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println(i * a)
	}
}

func main() {
	runtime.GOMAXPROCS(3)
	wg.Add(2)
	go printer(1)
	go printer(2)
	wg.Wait()
}
