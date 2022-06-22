package main

import "fmt"

func foo() {
	defer fmt.Println("foo Hello")

	fmt.Println("foo World")
}

func main() {
	foo()

	defer fmt.Println("Hello")

	fmt.Println("World")
}
