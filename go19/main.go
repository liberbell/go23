package main

import (
	"fmt"
	"os"
)

func foo() {
	defer fmt.Println("foo Hello")

	fmt.Println("foo World")
}

func main() {

	// defer fmt.Println("Hello")

	// foo()

	// fmt.Println("World")

	// fmt.Println("run")
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// fmt.Println("Success.")

	file, _ := os.Open("./lessoon.go")
	defer file.Close()

	data := make([]byte, 100)
}
