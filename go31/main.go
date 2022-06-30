package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	mike := Person{"Mile", 20}
	fmt.Println(mike)
}
