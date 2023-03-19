package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideArea float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {

	s := square{
		sideArea: 20,
	}
	t := triangle{
		height: 20,
		base:   5,
	}

	printArea(s)
	printArea(t)

}

func (sq square) getArea() float64 {
	return sq.sideArea * sq.sideArea
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
