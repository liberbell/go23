package main

import "fmt"

func do(i interface{}) {
	// ii := i.(int)
	// ii *= 2
	// fmt.Println(ii)
	ss := i.(string)
	fmt.Println(ss + "!")
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	case bool:
		fmt.Println(v)
	}
}

func main() {
	// var i interface{} = 10
	do(10)
	do("Mike")
	do(true)
}
