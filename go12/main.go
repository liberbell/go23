package main

import "fmt"

func add(x int, y int) (int, int) {
	fmt.Println("Add function")
	// fmt.Println(x + y)
	return x + y, x - y
}

func main() {
	result := add(10, 20)
	fmt.Println(result)
}
