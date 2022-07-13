package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

var s *semaphore.Weighted = semaphore.NewWeighted(1)

func lognProcess(ctx context.Context) {
	if err := s.Acquire(ctx, 1); err != nil {
		fmt.Println(err)
		return
	}
	defer s.Release(1)
	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done.")
}

func main() {
	ctx := context.TODO()
	go lognProcess(ctx)
	go lognProcess(ctx)
	go lognProcess(ctx)
	time.Sleep(5 * time.Second)
}
