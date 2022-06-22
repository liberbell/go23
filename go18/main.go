package main

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "mac"
}

func main() {

	// switch os := getOsName(); os {
	// case "mac":
	// 	fmt.Println("Mac!")
	// case "windows":
	// 	fmt.Println("Windows!")
	// default:
	// 	fmt.Println("I don`t know.", os)
	// }

	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning.")
	}

}
