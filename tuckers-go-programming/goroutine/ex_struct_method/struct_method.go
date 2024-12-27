package main

import "fmt"

type controller struct {
	command string
}

func (c *controller) editCommand(command string) {
	c.command = command
}

func main() {
	var c controller
	//c := controller{}

	c.editCommand("abc")
	fmt.Println(c.command)
}
