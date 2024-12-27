package main

import "fmt"

func main() {
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Bond":        2,
			"Hello":       4,
			"Fuckingthem": 6,
			"Me":          10,
		},
		favDrinks: []string{
			"H", "I", "J", "K", "L", "M", "N",
		},
	}
	fmt.Println(s.first)
}
