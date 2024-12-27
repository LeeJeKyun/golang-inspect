package main

import "fmt"

func main() {
	strings := []string{"abc", "bcd", "cde"}
	i := append(strings, "abc2")
	//stringptr := &strings
	//strings2 := *stringptr

	j := append(strings, "abc3")

	fmt.Println(strings)
	fmt.Println(i)
	fmt.Println(j)

	for i, n := range strings {
		fmt.Println(i, n)
	}
}
