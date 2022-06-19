package main

import "fmt"

func add(x int, y int) (int, int) {
	fmt.Println("Add function")
	// fmt.Println(x + y)
	return x + y, x - y
}

func cal(price, item int) (result int) {
	result = price * item
	return result
}

func converter(price int) float64 {
	return float64(price)
}

func main() {
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 := cal(100, 9)
	fmt.Println(r3)
}
