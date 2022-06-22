package main

import "fmt"

func foo() {
	defer fmt.Println("foo Hello")

	fmt.Println("foo World")
}

func main() {

	// defer fmt.Println("Hello")

	// foo()

	// fmt.Println("World")

	fmt.Println("run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("Success.")
}
