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
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	go lognProcess(ctx, ch)

CTXLOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break CTXLOOP
		case <-ch:
			fmt.Println("success")
			break CTXLOOP
		}
	}
	fmt.Println("End of Loop.")
}
