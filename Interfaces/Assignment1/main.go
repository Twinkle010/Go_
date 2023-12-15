package main

import (
	"fmt"
)

type shape interface{
	getArea() float64
}
type square struct{
	sideLength float64
}
type triangle struct{
	height float64
	base float64
}

func main() {
	s := square{}
	s.sideLength = 5
	t := triangle{
		base: 5,
		height: 10,
	}
	printArea(s)
	printArea(t)
}

func (s square) getArea() float64{
	return s.sideLength*s.sideLength
}
func (t triangle) getArea() float64{
	return 0.5* t.base * t.height
}

func printArea(sh shape) {
	fmt.Println("Area of the shape:", sh.getArea())
}