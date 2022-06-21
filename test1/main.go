package main

import (
	"fmt"
)

func main() {
	f := 1.11
	i := int(f)
	// conv_f, _ := strconv.Atoi(f)
	fmt.Println(i)

	m := map[string]int{
		"Mike":  20,
		"Nancy": 24,
		"Messi": 30,
	}
	fmt.Printf("%T %v\n", m, m)
}
