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

	m := map[string]person{
		a.last: a,
		b.last: b,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

}
