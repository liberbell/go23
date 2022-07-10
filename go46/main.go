package main

import (
	"context"
	"fmt"
	"time"
)

func lognProcess(ctx context.Context, ch chan string) {
	fmt.Println("Run")
	time.Sleep(2 * time.Second)
	fmt.Println("Finish")
	ch <- "result"
}

func main() {
	ch := make(chan string)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	go lognProcess(ctx, ch)

	for {
		select {
		case <-ch:
			fmt.Println("success")
			return
		}
	}
}
