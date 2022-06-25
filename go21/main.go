package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("./sample.txt")
	if err != nil {
		log.Fatalln(err)
	}
}
