package main

import (
	"github.com/aobelets/my-exemple-mod/shape"
)

func main() {

	square := shape.Square{SideLenght: 5}
	circle := shape.Circle{Radius: 8}

	shape.PrintShapeArea(square)
	shape.PrintShapeArea(circle)
}
