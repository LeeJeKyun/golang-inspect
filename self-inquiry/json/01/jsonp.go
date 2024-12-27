package main

import (
	"encoding/json"
	"fmt"
)

type hello struct {
	Name  string
	Check bool
	Dday  string
}

func main() {
	h := hello{"오늘 할 일", false, "20241218"}

	marshal, _ := json.Marshal(h)
	//if err != nil {
	//	fmt.Println("에러발생")
	//	return
	//}

	jsonString := string(marshal)
	fmt.Println(jsonString)

	//file, err := os.Create("/Users/jekyunlee/json/hello.json")
	//if err != nil {
	//	fmt.Println("파일생성 에러 발생")
	//	return
	//}
	//
	//defer file.Close()
	//
	//file.WriteString(jsonString)
	//
	//file.Sync()

}
