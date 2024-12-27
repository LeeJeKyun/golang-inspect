package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
	Desc  string
}

func main() {
	p := person{"Lee", "Jekyun", 30, "Hello, Golang"}
	marshalledP, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	//os.Stdout.Write(marshalledP)
	fmt.Println(string(marshalledP))

	var p2 person

	err2 := json.Unmarshal(marshalledP, &p2)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("all of the data", p2)

}
