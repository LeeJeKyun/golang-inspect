package main

import "fmt"

//&: 주소를 전달한다.
//var a *int = *i
//*: 주소를 역참조하거나 특정 위치 또는 특정 주소에 저장된 값을 가져온다.
//var b int = *b

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)
	//int Type
	fmt.Printf("%T\n", a)
	//*int Type
	fmt.Printf("%T\n", &a) // & 는 주소를 전달한다.
	//int, *int
	//int와 int에 대한 포인터는 별개의 두가지 타입이다.

	b := &a
	fmt.Println(b)
	fmt.Println(*b) // *는 주소에 저장된 값을 전달한다.
	fmt.Println(*&a)

	// b주소에 저장된 값을 43으로 대입한다.
	*b = 43
	fmt.Println(a)
}
