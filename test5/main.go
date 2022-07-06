package main

import "fmt"

func goroutine(s []string, c chan int) {
	sum := ""
	for _, v := range s {
		sum += v
		c <- sum
	}
}

func main() {
	words := []string{"test1!", "test2!", "test3!", "test4!"}
	c := make(chan int)
	go goroutine(words, c)
	for w := range c {
		fmt.Println(w)
	}
}
