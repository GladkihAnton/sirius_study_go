package interfaces

import "math"

type geometry interface {
	Area() float32
}

type Square struct {
	width  int
	height int
}

func (square Square) Area() float32 {
	return float32(square.width * square.height)
}

type Circle struct {
	radius float32
}

func (square Circle) Area() float32 {
	return square.radius * square.radius * math.Phi
}

func calcArea(geo geometry) {

}

func Main() {
	MainTask1()
	//square := Square{1, 2}
	//circle := Circle{1}
	//calcArea(square)
	//calcArea(circle)
	//MainEmb()
}
