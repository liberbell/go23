package main

import "fmt"

func incrementgenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {

	fmt.Println(incrementgenerator())
	fmt.Println(incrementgenerator())
	fmt.Println(incrementgenerator())
}
