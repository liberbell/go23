package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("Continue")
			continue
		}
		fmt.Println(i)
	}
}
