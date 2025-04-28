package main

import "fmt"

func main() {

	//오름차순으로 고정
	s := []int{5, 22, 54, 23, 4, 11, 100, 143, 30, 36, 12, 18}

	selectionSort(&s)

	fmt.Println(s)

}

func bubbleSort(s *[]int) {
	for i := 0; i < len(*s)-1; i++ {
		for j := i; j < len(*s); j++ {
			if (*s)[i] > (*s)[j] {
				(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
			}
		}
	}
}

func selectionSort(s *[]int) {
	for i := 0; i < len(*s)-1; i++ {
		min := i
		for j := i + 1; j < len(*s); j++ {
			if (*s)[j] < (*s)[min] {
				min = j
			}
		}
		(*s)[i], (*s)[min] = (*s)[min], (*s)[i]
	}
}

func insertionSort(s *[]int) {
	sp := *s
	for i := 1; i < len(sp); i++ {
		temp := (sp)[i]
		prev := i - 1
		for (prev >= 0) && sp[prev] > temp {
			sp[prev+1] = sp[prev]
			prev--
		}
	}
}
