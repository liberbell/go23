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

func main() {
	ctx := context.TODO()
	go lognProcess(ctx)
	go lognProcess(ctx)
	go lognProcess(ctx)
	time.Sleep(5 * time.Second)
}
