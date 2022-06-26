package main

import "fmt"

func main() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	i := 0
	for _, value := range l {
		// fmt.Println(l[i])
		if l[i] >= value {
			fmt.Println(l[i], value)
			result := value
			fmt.Println("Value:", result)
		}
		if i < len(l) {
			i = i + 1
			// fmt.Println(l[i])
		} else {
			fmt.Println("Result: ", l[i])
		}
	}
	// for _, value1 := range l {
	// 	num1 := value1

	// 	switch value1 >  {
	// 	case condition:

	// 	}
	// }
}
