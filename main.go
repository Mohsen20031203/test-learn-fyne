package main

import "math"

type Shape interface {
	Area() float64
	Numebr() bool
}

type Circle struct {
	Radius float64
}

func newTest() Shape {
	return &Circle{}
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Numebr() bool {
	return false
}

type example struct {
	Number2 int
}

func newTestTwo() Shape {
	return &example{}
}

func (e *example) Area() float64 {
	return 0
}

func (e *example) Numebr() bool {
	return false
}

func main() {
	m := newTest()
	m.Area()
	m.Numebr()

	n := newTestTwo()
	n.Area()
}
