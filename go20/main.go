package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatalln("Exit", err)
	}

	log.Println("logging!")
	log.Printf("%T %v", "test", "test")

	log.Fatalf("%T %v", "test", "test")
	log.Fatalln("Error!")
	fmt.Println("Ok.")
}
