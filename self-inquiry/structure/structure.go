package main

import "fmt"

type User struct {
	Name string
	id   int
	age  int
}

func main() {
	var user User
	user.Name = "Hello"
	user.id = 1
	user.age = 20
	fmt.Println(user)
}
