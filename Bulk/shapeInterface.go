package main

import (
	"fmt"
	"math"
)

// SHAPE INTERFACE
type shape interface {
	area() float64
	circumf() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type rect struct {
	length, breadth float64
}

//SQUARE METHODS
func (s square) area() float64 {
	return s.length * s.length
}
func (s square) circumf() float64 {
	return s.length * 4
}


//CIRCLE METHODS
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

//RECTANGLE METHODS
func (r rect) area() float64 {
	return r.length * r.breadth
}
func (r rect) circumf() float64 {
	return 2 * (r.length + r.breadth)
}


func printShapeInfo(s shape) {
	fmt.Printf("Area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("Circumference of %T is: %0.2f \n", s, s.circumf())
}

func main() {
	shapes := []shape{
		square{length: 15.2},
		circle{radius: 7.5},
		rect{length: 5.5, breadth: 10.5},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("---")
	}
}