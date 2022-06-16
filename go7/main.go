package main

import "fmt"

func main() {
	fmt.Println("Hello World.")
	fmt.Println("Hello " + "World.")
	fmt.Println(string("Hello World."[0]))

	var s string = "Hello World."
	s[0] = "A"

}
