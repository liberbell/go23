package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

func (r Vertex) Plus() int {
	return r.X + r.Y
}

func main() {
	// v := Vertex{3, 4}
	// fmt.Println(v.Plus())
	v := Vertex{3, 4}
	fmt.Println(v)
}
