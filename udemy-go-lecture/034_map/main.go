package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	fmt.Println(m["James"])

	fmt.Println(m["Barnaba"])

	v, ok := m["Barnaba"]
	fmt.Println(v)
	fmt.Println(ok)

	m["todd"] = 33
	fmt.Println(m)

	if v, ok := m["Barnabas"]; ok {
		fmt.Println(v)
	}
	for k, v := range m {
		fmt.Println("key=", k, "value=", v)
	}

	xi := []int{4, 5, 7, 8, 9, 42}

	for i, v := range xi {
		fmt.Println(i, v)
	}

	fmt.Println()

	yi := make([]int, 6, 100)
	yi[0] = 210
	yi[1] = 211
	yi[2] = 212
	yi[3] = 213
	yi[4] = 214
	yi[5] = 215

	for i, v := range yi {
		fmt.Println(i, v)
	}

}
