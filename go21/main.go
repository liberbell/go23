package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("./sample.txt")
	if err != nil {
		log.Fatal("Error", err)
	}
	defer file.Close()
	data := make([]byte, 0)
}
