package main

func producer(ch chan int, i int) {
	ch <- i
}

func main() {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}
}
