package main

import "fmt"

func main() {
	x := 1
	y := 2
	fmt.Printf("adding %v and %v\n", x, y)
	fmt.Println("result from operations: ", Add(x, y))
}

func Add(x, y int) int {
	res := x + y
	res += 1
	return res
}
