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

func (v Vertex) String() string {
	return fmt.Sprintf("X is %v! Y is %v!\n", e.X, e.Y)
}

func main() {
	// v := Vertex{3, 4}
	// fmt.Println(v.Plus())
	v := Vertex{3, 4}
	fmt.Println(v)
	// e1 := &ValuePrint{3, 4}
}
