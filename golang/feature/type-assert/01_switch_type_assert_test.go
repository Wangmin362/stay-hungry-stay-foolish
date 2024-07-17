package type_assert

import (
	"fmt"
	"testing"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func TestSwitchAssert(t *testing.T) {
	shapes := []Shape{
		Circle{Radius: 2},
		Rectangle{Width: 3, Height: 4},
		Circle{Radius: 5},
	}

	for _, shape := range shapes {
		switch shape := shape.(type) {
		case Circle:
			fmt.Printf("Circle: Area=%.2f\n", shape.Area())
		case Rectangle:
			fmt.Printf("Rectangle: Area=%.2f\n", shape.Area())
		default:
			fmt.Println("Unknown shape")
		}
	}
}
