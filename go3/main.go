package main

import (
	"fmt"
	"os/user"
	"time"
)

func init() {
	fmt.Println("init")
}

func bazz() {
	fmt.Println("Bazz")
}

func main() {
	fmt.Println("Hello world", "Test Test", time.Now())
	fmt.Println(user.Current())

	// bazz()
}
