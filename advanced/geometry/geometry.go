package geometry

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Height + 2*r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func printShapeInfo(shapeName string, s Shape) {
	area := s.Area()
	perimeter := s.Perimeter()

	fmt.Println(shapeName)
	fmt.Printf("Area : %.2f\n", area)
	fmt.Printf("Perimeter : %.2f\n", perimeter)

}
