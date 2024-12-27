package main

import "fmt"

type People struct {
	name string
	age  int
}

type Man struct {
	name string
	age  int
}

// 타입변환과 관련된 탐구
func main() {
	i := 100
	u := uint(i)
	f := float32(u)

	fmt.Println(f)
	fmt.Println(i)

	str := "ABC"
	var bytes []byte = []byte(str)
	fmt.Println("String:", str)
	fmt.Println("stringToBytes:", bytes)

	bytesToStr := string(bytes)
	fmt.Println("bytesToStr:", bytesToStr)

	p := new(People)
	fmt.Println("people:", p)
	p.name = "jekyun"
	p.age = 31
	fmt.Println("people:", *p)
}
