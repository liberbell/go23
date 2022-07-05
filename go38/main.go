package main

import (
	"fmt"
	"time"
)

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

func goroutine2(ch chan int) {
	for {
		ch <- 100
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
	c2 := make(chan int)

	go goroutine1(c1)
	go goroutine2(c2)

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
