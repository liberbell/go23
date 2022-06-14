package main

import "fmt"

func init() {
	fmt.Println("init")
}

func bazz() {
	fmt.Println("Bazz")
}

func main() {
	fmt.Println("Hello world")

	bazz()
}
