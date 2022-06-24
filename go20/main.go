package main

import (
	"fmt"
	"log"
	"os"
)

func loggingSetting(logFile string) {
	logFile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
}

func main() {
	_, err := os.Open("sample.txt")
	if err != nil {
		log.Fatalln("Exit", err)
	}

	log.Println("logging!")
	log.Printf("%T %v", "test", "test")

	log.Fatalf("%T %v", "test", "test")
	log.Fatalln("Error!")
	fmt.Println("Ok.")
}
