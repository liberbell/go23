package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	// return "My name is: " + p.Name
	return fmt.Sprintf("My name is %v.", p.Name)
}

func main() {
	mike := Person{"Mike", 20}
	fmt.Println(mike)
}
