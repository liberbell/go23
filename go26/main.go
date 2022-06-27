package main

type Vertex struct {
	X int
	Y int
	S string
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	v.X = 1000
}

func main() {
	// v := Vertex{X: 10, Y: 20}
	// fmt.Println(v)
	// fmt.Println(v.X, v.Y)

	// v.X = 1000
	// fmt.Println(v)

	// v2 := Vertex{X: 200}
	// fmt.Println(v2)

	// v3 := Vertex{4, 8, "test"}
	// fmt.Println(v3)

	// v4 := Vertex{}
	// fmt.Println(v4)
	// fmt.Printf("%T\n", v4)

	// var v5 Vertex
	// fmt.Println(v5)
	// fmt.Printf("%T\n", v5)

	// v6 := new(Vertex)
	// fmt.Println(v6)
	// fmt.Printf("%T\n", v6)

	// v7 := &Vertex{}
	// fmt.Printf("%T %v\n", v7, v7)

	// s := []int{}
	// fmt.Println(s)

}
