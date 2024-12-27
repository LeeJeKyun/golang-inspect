package main

import "fmt"

type Hello struct {
	Name  string
	Check bool
}

func main() {
	hellos := []Hello{{"A", false}, {"B", false}, {"C", false}}

	hellosp := &hellos

	hellosC := *hellosp

	for i, n := range hellosC {
		if "A" == n.Name {
			hellosC[i].Check = true
		}
	}
	fmt.Println(hellosC)
}
