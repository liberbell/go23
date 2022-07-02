package main

import "fmt"

func myFunc() error {
	return nil
}

func main() {
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}
