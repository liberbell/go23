package main

import "fmt"

func main() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	fmt.Println(m)
	m["new"] = 400
	fmt.Println(m)

	fmt.Println(m["noitem"])

	v, ok := m["noitem"]
	fmt.Println(v, ok)

	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)

	// var m3 map[string]int
	// m3["pc"] = 5000
	// fmt.Println(m3)

	var s []int
	if s == nil {
		fmt.Println("Nil")
	}
}
