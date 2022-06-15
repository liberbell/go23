package main

import "fmt"

var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false
)

func foo() {
	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false

	fmt.Println(xi, xf64, xs, xt, xf)
}

func main() {
	// var (
	// 	i    int     = 1
	// 	f64  float64 = 1.2
	// 	s    string  = "test"
	// 	t, f bool    = true, false
	// )

	// var f bool = false

	fmt.Println(i, f64, s, t, f)
	fmt.Printf("%T", f64)

	foo()
}
