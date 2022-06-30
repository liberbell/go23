package main

import "fmt"

func do(i interface{}) {
	// ii := i.(int)
	// ii *= 2
	// fmt.Println(ii)
	ss := i.(string)
	fmt.Println(ss + "!")
}

func main() {
	// var i interface{} = 10
	// do(i)
	do("Mike")
	// do(true)
}
