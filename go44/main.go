package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{5, 3, 2, 8, 7}
	s := []string{"d", "a", "f"}
	p := []struct {
		Name string
		Age  int
	}{
		{"Bob", 43},
		{"Eric", 73},
		{"Alex", 34},
		{"Elton", 69},
	}
	fmt.Println(i, s, p)
	sort.Ints(i)
	fmt.Println(i)
}
