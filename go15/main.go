package main

import "fmt"

func by2(num int) string {
	if num%2 == 0 {
		return "Yes"
	} else {
		return "No"
	}
}

func main() {
	// num := 9
	// if num%2 == 0 {
	// 	fmt.Println("by 2")
	// } else if num%3 == 0 {
	// 	fmt.Println("by 3")
	// } else {
	// 	fmt.Println("else")
	// }
	// x, y := 11, 12
	// if x == 10 && y == 10 {
	// 	fmt.Println("&&")
	// }
	// if x == 10 || y == 10 {
	// 	fmt.Println("||")
	// }

	result := by2(10)
	if result == "Yes" {
		fmt.Println("result: ", result)
	}
}
