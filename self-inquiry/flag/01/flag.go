package main

import (
	"flag"
	"fmt"
)

func main() {

	//arg := os.Args[1]
	stringPtr := flag.String("test", "abc", "a string")

	flag.Parse()

	//fmt.Println("Arg:", arg)
	fmt.Println("Flag(test)", *stringPtr)

}
