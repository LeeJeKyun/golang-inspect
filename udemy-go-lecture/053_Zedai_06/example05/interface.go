package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return c.radius * math.Pi * c.radius
}

type square struct {
	length float64
}

func (s square) area() float64 {
	return s.length * s.length
}

type Shape interface {
	area() float64
}

func info(shape Shape) {
	fmt.Println(shape.area())
}

func main() {
	c1 := circle{123}
	info(c1)

	s1 := square{5}
	info(s1)
}
