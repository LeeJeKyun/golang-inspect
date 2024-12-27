package main

import "fmt"

func main() {
	switch "Bond" {
	case "Money", "bond", "Do No":
		fmt.Println("miss money or bond or do no")
	case "M":
		fmt.Println("M")
	case "Q":
		fmt.Println("this is Q")
	default:
		fmt.Println("default")
	}
}
