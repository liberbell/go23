package main

func main() {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}
}
