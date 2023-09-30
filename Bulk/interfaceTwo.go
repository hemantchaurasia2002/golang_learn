package main

type shape interface {
	circle int;
	rectange int;
}

type circle struct {
	radius int
}

type rectangle struct {
	width, height int
}

func (c circle) area() int {
	return math.pi * c.radius * c.radius
}

func (r rectangle) area() int {
	return r.width * r.height
}

func printarea(s shape) int {
	return s.area
}

func main () {
	Rectangle := rectangle{width: 10, height: 15}
	Circle := circle{radius: 10}
}