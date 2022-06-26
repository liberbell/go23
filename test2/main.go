package main

import "fmt"

func main() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	// i := 0
	result := 0
	for index, value := range l {
		// fmt.Println(l[i])
		fmt.Println(index, value)
		if index == 0 {
			result = value
			fmt.Println(index, value)
		} else if result < value {
			fmt.Println("p1:", value)
			result = value
		}
		fmt.Println("Result:", result)
		// 	if l[i] >= value {
		// 		fmt.Println(l[i], value)
		// 		result = value
		// 		fmt.Println("Value:", result)
		// 	}
		// 	fmt.Println("lengs: ", len(l))
		// 	fmt.Println("index: ", i)
		// 	if i < len(l) {
		// 		i = i + 1
		// 		// fmt.Println(l[i])
		// 	} else {

		// 	}
		// 	fmt.Println("Result: ", result)
		// }
		// for _, value1 := range l {
		// 	num1 := value1

		// 	switch value1 >  {
		// 	case condition:

		// 	}
		// }
	}
}
