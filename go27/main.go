package main

import "fmt"

type Vertex struct {
	x, Y int
}

func Area(v Vertex) int {
	return v.x * v.Y
}

func (v Vertex) Area() int {
	return v.x * v.Y
}

func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.Y = v.Y * i
}

func New(x, y int) *Vertex {
	return &Vertex{x, y}
}

func main() {
	// v := Vertex{3, 4}
	// fmt.Println(Area(v))
	v := New(3, 4)
	v.Scale(10)
	fmt.Println(v.Area())
}
