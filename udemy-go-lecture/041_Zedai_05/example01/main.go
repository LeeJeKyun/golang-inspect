package main

import "fmt"

func main() {
	type person struct {
		first      string
		last       string
		favFalvors []string
	}

	var a person = person{
		first:      "Lee",
		last:       "JeKyun",
		favFalvors: []string{"strawberry", "chocolate"},
	}

	var b person = person{
		first:      "Lee",
		last:       "YJ",
		favFalvors: []string{"water", "ice"},
	}

	fmt.Println(a.first)
	fmt.Println(a.last)
	for i, v := range a.favFalvors {
		fmt.Println(i, v)
	}

	fmt.Println(b.first)
	fmt.Println(b.last)
	for i, v := range b.favFalvors {
		fmt.Println(i, v)
	}

}
