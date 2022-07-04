package main

import (
	"fmt"
	"sync"
)

func producer(ch chan int, i int) {
	ch <- i * 2
}

func consumper(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Println("process:", i*1000)
		wg.Done()
	}
	fmt.Println("#########")
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}
	go consumper(ch, &wg)
	wg.Wait()
	close(ch)
}
