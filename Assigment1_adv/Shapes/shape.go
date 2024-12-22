package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimetr() float64
}
type Restangle struct {
	Length float64
	Width  float64
}
type Circle struct {
	Radius float64
}
type Square struct {
	Length float64
}
type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

func (r Restangle) Area() float64 {
	return r.Length * r.Width

}
func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}
func (s Square) Area() float64 {
	return s.Length * s.Length
}
func (t Triangle) Area() float64 {
	p := (t.SideA + t.SideB + t.SideC) / 2
	return math.Sqrt(p * (p - t.SideA) * (p - t.SideB) * (p - t.SideC))
}

func (r Restangle) Perimetr() float64 {
	return (r.Length + r.Width) * 2
}
func (c Circle) Perimetr() float64 {
	return 2 * math.Pi * c.Radius
}
func (s Square) Perimetr() float64 {
	return 4 * s.Length
}
func (t Triangle) Perimetr() float64 {
	return t.SideA + t.SideB + t.SideC
}
func PrintShapeDetails(s Shape) {
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimetr:", s.Perimetr())
}
func main() {

	shapes := []Shape{
		Restangle{Length: 10, Width: 5},
		Circle{Radius: 7},
		Square{Length: 4},
		Triangle{SideA: 3, SideB: 4, SideC: 5},
	}

	for _, shape := range shapes {
		fmt.Println("---------------------")
		PrintShapeDetails(shape)
	}
}
