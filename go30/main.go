package main

import "fmt"

func do(i interface{}) {
	// ii := i.(int)
	// ii *= 2
	// fmt.Println(ii)
	// ss := i.(string)
	// fmt.Println(ss + "!")
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	case bool:
		fmt.Printf("unknown type Type:%T Value:%v\n", v, v)
	}
}

func main() {
	// var i interface{} = 10
	do(10)
	do("Mike")
	do(true)

	var i int = 10
	ii := float64(10)
	fmt.Println(i, ii)
}
