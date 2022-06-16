package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World.")
	fmt.Println("Hello " + "World.")
	fmt.Println(string("Hello World."[0]))

	var s string = "Hello World."

	fmt.Println(strings.Replace(s, "H", "A", 1))
	fmt.Println(strings.Contains(s, "World"))

	fmt.Println("Test" +
		"Test")
}
