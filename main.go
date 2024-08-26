/*
	1. chenge struct project interface's
	2. Project layering's
	3. Project folder's
*/

package main

import (
	"fmt"
	"math"
)

// تعریف یک interface به نام Shape
type Shape interface {
	Area() float64
}

// پیاده‌سازی برای Rectangle
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// پیاده‌سازی برای Circle
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// پیاده‌سازی برای Triangle
type Triangle struct {
	Base, Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// تابع عمومی برای چاپ مساحت
func printArea(shape Shape) {
	fmt.Println("Area:", shape.Area())
}

func main() {
	r := Rectangle{Width: 10, Height: 5}
	c := Circle{Radius: 7}
	c1 := Circle{Radius: 10}
	t := Triangle{Base: 10, Height: 8}

	printArea(r)
	printArea(c)
	printArea(c1)
	printArea(t)
}
