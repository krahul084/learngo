package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	baseLength float64
	height     float64
}

type square struct {
	side float64
}

func main() {
	nt := triangle{}

	ns := square{}
	nt.baseLength = 4
	nt.height = 4.0
	ns.side = 3.7

	printArea(nt)
	printArea(ns)
}

func (t triangle) getArea() float64 {
	return (0.5 * t.baseLength * t.height)
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
