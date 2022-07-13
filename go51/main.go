package main

import (
	"context"
	"fmt"
	"time"
)

func lognProcess(ctx context.Context) {
	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done.")
}
