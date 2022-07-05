package main

import "time"

func goroutine(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func goroutine1(ch chan string) {
	for {
		ch <- "packet from 1"
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	// s := []int{1, 2, 3, 4, 5}
	// c := make(chan int)
	// go goroutine(s, c)
	// go goroutine(s, c)

	// x := <-c
	// fmt.Println(x)
	// y := <-c
	// fmt.Println(y)
	c1 := make(chan string)
	c2 := make(chan string)

	go goroutine()
}
