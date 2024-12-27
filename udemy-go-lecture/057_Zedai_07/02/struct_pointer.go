package main

import "fmt"

type person struct {
	first string
}

func changeMe(p1 person) {
	//p1.first = "Miss Moneypenny"
	(&p1).first = "Miss Moneyp"
}

func main() {
	p1 := person{
		first: "James Bond",
	}
	fmt.Println(p1)
	changeMe(p1)
	fmt.Println(p1)
}
