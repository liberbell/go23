package main

import "fmt"

func main() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	for _, v := range l {
		fmt.Println(v)
	}
}
