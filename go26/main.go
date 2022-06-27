package main

import "fmt"

type Vertex struct {
	X int
	Y int
	S string
}

func main() {
	v := Vertex{X: 10, Y: 20}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)

	v.X = 1000
	fmt.Println(v)

	v2 := Vertex{X: 200}
	fmt.Println(v2)

	v3 := Vertex{4, 8, "test"}
	fmt.Println(v3)

	v4 := Vertex{}
	fmt.Println(v4)

	var v5 Vertex
	fmt.Println(v5)
}
