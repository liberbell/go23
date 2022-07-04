package main

import "fmt"

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
}

func goroutine2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int, len(s))
	go goroutine1(s, c)
	// go goroutine2(s, c)
	// x := <-c
	// fmt.Println(x)
	// y := <-c
	// fmt.Println(y)

	for i := range c {
		fmt.Println(i)
	}
}
