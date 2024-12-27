package main

import "fmt"

func main() {
	ch := make(chan int)

	ch <- 9
	//채널에 데이터를 넣어두고 빼가지 않아서 데드락 발생
	fmt.Println("Never Print")
}
