package main

import (
	"fmt"
	"sort"
)

func main() {
	ans := 0

	var n, m int
	fmt.Scan(&n, &m)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	//오름차순 정렬
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	sp := 0
	ep := len(arr) - 1

	for sp < ep {
		if arr[sp]+arr[ep] > m {
			ep--
		} else if arr[sp]+arr[ep] < m {
			sp++
		} else if arr[sp]+arr[ep] == m {
			sp++
			ep--
			ans++
		}
	}
	fmt.Println(ans)

}
