package main

import "fmt"

func main() {
	arr := [5]*int{0: new(int), 3: new(int)}

	*arr[0] = 10
	//*arr[1] = 20

	fmt.Println(*arr[0])
	//fmt.Println(*arr[1])
	fmt.Println(*arr[3])
}
