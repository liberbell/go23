package main

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)
}
