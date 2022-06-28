package main

type Vertex struct {
	X, Y int
}

func Area(v Vertex) int {
	return v.X * v.Y
}

func main() {
	v := Vertex{3, 4}
}
