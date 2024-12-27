package main

import (
	"encoding/json"
	"fmt"
)

type hello struct {
	Name string
	Key  bool
}

func main() {
	h1 := hello{"A", true}
	h2 := hello{"B", true}
	h3 := hello{"C", false}
	h4 := hello{"D", false}

	hs := []hello{h1, h2, h3, h4}

	bytes, _ := json.Marshal(hs)
	jsonString := string(bytes)

	fmt.Println("jsonString:", jsonString)

	unBytes := []byte(jsonString)

	unHs := []hello{}

	json.Unmarshal(unBytes, &unHs)

	fmt.Println(unHs)

}
