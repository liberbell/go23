package main

import (
	"fmt"
	"time"
)

func lognProcess(ch chan string) {
	fmt.Println("Run")
	time.Sleep(2 * time.Second)
	fmt.Println("Finish")
	ch <- "result"
}

func main() {
	ch := make(chan string)
	go lognProcess(ch)
}
