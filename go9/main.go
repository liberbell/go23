package main

import "fmt"

func main() {
	// var a [2]int
	// a[0] = 100
	// a[1] = 200

	// fmt.Println(a)

	// var b [2]int = [2]int{100, 200}

	// fmt.Println(b)

	// var c []int = []int{100, 200}
	// c = append(c, 300)
	// fmt.Println(c)

	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	fmt.Println(n[2])
	fmt.Println(n[2:4])
	fmt.Println(n[:2])
	fmt.Println(n[2:])
	fmt.Println(n[:])

	n[2] = 100
	fmt.Println(n)

	var board = [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}
	fmt.Println(board)

	n = append(n, 100, 200, 300, 400)
	fmt.Println(n)
}
