package shape

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
}

type Square struct {
	SideLenght float32
}

func (s Square) Area() float32 {
	return s.SideLenght * s.SideLenght
}

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return c.Radius * c.Radius * math.Pi
}

func PrintShapeArea(shape Shape) {
	fmt.Println(shape.Area())

}
