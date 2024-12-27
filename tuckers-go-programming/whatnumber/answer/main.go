package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var _ rand.Source
	_ = rand.NewSource(time.Now().UnixNano())

	n := rand.Intn(100)

	fmt.Println(n)
}
