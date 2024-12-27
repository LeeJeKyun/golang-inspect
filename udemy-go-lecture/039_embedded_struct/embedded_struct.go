package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk   bool
	first string
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk:   true,
		first: "Dummy",
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
	}
	sa1.age = 42
	p2.age = 30

	fmt.Println(sa1)
	fmt.Println(p2)

	fmt.Println(sa1.first, sa1.person.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println(p2.first, p2.last, p2.age)

}
