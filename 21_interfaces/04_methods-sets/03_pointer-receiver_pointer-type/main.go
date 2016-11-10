package main

import (
	"math"
	"fmt"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("Area", s.area())
}

func main() {
	c := circle{5}
	info(&c)
}


// pointer receiver has to get a pointer type