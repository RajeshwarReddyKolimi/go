package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	length float64
	width  float64
}

func (r rectangle) area() float64 {
	return r.length * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func Exercise3() {
	c := circle{radius: 10}
	r := rectangle{length: 10, width: 20}

	fmt.Println("Area of circle:", c.area())
	fmt.Println("Area of rectangle:", r.area())

	shapes := []shape{c, r}
	for _, shape := range shapes {
		fmt.Println("Area:", shape.area())
	}
}
