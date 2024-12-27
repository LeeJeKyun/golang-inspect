package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p *person) speak() {
	fmt.Println("My Name is", p.first, p.last, "And My Age is", p.age)
}

func main() {
	p := person{"Lee", "Jekyun", 30}
	p.speak()
}
