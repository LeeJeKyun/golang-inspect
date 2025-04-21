package main

import (
	"fmt"
	"sync"
)

type selfMutex struct {
	lock bool
}

func (m *selfMutex) Lock() {
	for m.lock == true {
	}
	if m.lock == false {
		m.lock = true
	}
}

func (m *selfMutex) Unlock() {
	m.lock = false
}

var a int

var wg sync.WaitGroup

func main() {
	var m selfMutex
	wg.Add(2)
	go increment(&m)
	go increment(&m)
	wg.Wait()
	fmt.Println(a)
}
func increment(m *selfMutex) {
	//m.Lock()
	for i := 0; i < 10000; i++ {
		a++
	}
	//m.Unlock()
	wg.Done()
}
