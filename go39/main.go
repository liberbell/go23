package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case t := <-tick:
			fmt.Println("Tick!", t)
		case <-boom:
			fmt.Println("Boom!")
			return
		default:
			fmt.Println("   .")
			time.Sleep(50 * time.Millisecond)
		}
	}
	fmt.Println("###########")
}
