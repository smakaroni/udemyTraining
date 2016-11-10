package main

import (
	"math"
	"fmt"
)

type Square struct {
	side float64
	test string
}

type Circle struct {
	radius float64
	test string
}

type Shape interface {
	area() float64
	ttest() string
}

func (s Square) area() float64  {
	return s.side * s.side
}
func (s Square) ttest() string {
	return "FU"
}

func (c Circle) area() float64  {
	return math.Pi * c.radius * c.radius
}
func (c Circle) ttest() string {
	return "FU"
}

func info(z Shape)  {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := Square{20, "t"}
	c := Circle{10, "t"}
	info(s)
	info(c)
}
