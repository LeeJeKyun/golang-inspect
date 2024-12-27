package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

// circle객체 메소드의 리시버가 circle의 포인터일 경우 주소 값(&)으로 실행이 가능하다.
// circle객체 메소드의 리시버가 circle일 경우 값과 주소값 모두로 실행이 가능하다.
func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{5}
	info(&c)
	//info(c) //현재 리시버가 포인터이므로 circle 타입으로 실행 불가
}
