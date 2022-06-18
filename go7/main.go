package main

import (
	"fmt"
)

func main() {
	// 	fmt.Println("Hello World.")
	// 	fmt.Println("Hello " + "World.")
	// 	fmt.Println(string("Hello World."[0]))

	// 	var s string = "Hello World."

	// 	fmt.Println(strings.Replace(s, "H", "A", 1))
	// 	fmt.Println(strings.Contains(s, "World"))

	// 	fmt.Println(`Test
	// Test`)
	// 	fmt.Println("\"")
	// 	fmt.Println(`"`)

	// var t, f bool = true, false
	// fmt.Println(t, f)
	// t, f := true, false
	// fmt.Println(t, f)
	// fmt.Printf("%T %v %t\n", t, t, t)
	// fmt.Printf("%T %v %t\n", f, f, f)

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println((false && false))

	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)

	fmt.Println(!true)
	fmt.Println(!false)
}
