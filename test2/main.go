package main

func main() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	i := 0
	for _, value := range l {
		switch {
		case l[i] > value:
			result := value
			i = i + 1
		}
	}
	// for _, value1 := range l {
	// 	num1 := value1

	// 	switch value1 >  {
	// 	case condition:

	// 	}
	// }
}
