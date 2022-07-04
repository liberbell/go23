package main

func producer(first chan int) {
	defer close(first)
	for i := 0; i < 10; i++ {
		first <- i
	}
}

func main() {
	first := make(chan int)
	second := make(chan int)
	third := make(chan int)

}
