package main

import (
	"fmt"
	"time"
)

func lognProcess(ch chan string) {
	fmt.Println("Run")
	time.Sleep(2 * time.Second)
}
