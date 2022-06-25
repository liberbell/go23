package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./sample.txt")
	if err != nil {
		log.Fatalln("Error", err)
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error", err)
	}
	fmt.Println(count, string(data))

	err = os.Chdir("test")
	if err != nil {
		log.Fatalln("Error!", err)
	}
}
