package main

import (
	"fmt"
	"sync"
)

type person struct {
	num  int
	name string
	age  int
}

var wg sync.WaitGroup

func service(c chan person) {
	for msg := range c {
		fmt.Println(msg)
	}
}

func main() {
	wg.Add(1)
	c := make(chan person)
	var num int
	var age int
	var name string
	var p person
	go service(c)
	for {
		fmt.Scanf("%d %s %d", &num, &name, &age)
		if num == 0 {
			break
			close(c)
			wg.Done()
		}
		p = person{num, name, age}
		c <- p
	}
}
