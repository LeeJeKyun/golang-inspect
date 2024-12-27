package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type hello struct {
	Name  string
	Check bool
	Dday  string
}

func main() {
	helloArray := []hello{{"A", false, "20241201"}, {"B", true, "20241219"}}

	marshal, err := json.Marshal(helloArray)
	if err != nil {
		fmt.Println("Json.Marshal Fail")
		return
	}

	jsonString := string(marshal)
	fmt.Println(jsonString)

	file, err := os.Create("/Users/jekyunlee/json/helloArray.json")
	if err != nil {
		fmt.Println("File Create Fail")
		return
	}
	defer file.Close()

	file.Write(marshal)
}
