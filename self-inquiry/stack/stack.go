package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	v *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (q *Stack) Push(v interface{}) {
	q.v.PushBack(v)
}

func (q *Stack) Pop() interface{} {
	back := q.v.Back()
	if back == nil {
		return nil
	}

	return q.v.Remove(back)
}

func main() {
	s := "()(()()(()))"
	solution(s)
}

func solution(s string) bool {
	bytes := []byte(s)

	for _, v := range bytes {
		sprintf := fmt.Sprintf("%c", v)
		fmt.Println(sprintf)
	}

	return false
}
