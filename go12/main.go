package main

import "fmt"

func add(x int, y int) (int, int) {
	fmt.Println("Add function")
	// fmt.Println(x + y)
	return x + y, x - y
}

func main() {
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)
}
