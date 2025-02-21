package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string `tag1:"이름" tag2:"Name"`
	age  int    `tag1:"나이" tag2:"Age"`
}

func main() {
	p := Person{}

	name, ok1 := reflect.TypeOf(p).FieldByName("name")
	fmt.Println(ok1, name.Tag.Get("tag1"), name.Tag.Get("tag2"))

	age, ok2 := reflect.TypeOf(p).FieldByName("age")
	fmt.Println(ok2, age.Tag.Get("tag1"), age.Tag.Get("tag2"))
}
