package main

import "time"

func normal(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
	}
}
