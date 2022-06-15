package main

import "fmt"

const pi = 3.14

const (
	Username = "testuser"
	Password = "password"
)

var big int = 9223372036854775807

func main() {
	fmt.Println(pi, Username, Password)
	fmt.Println(big)
}
