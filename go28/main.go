package main

import "fmt"

type Myint int

func (i Myint) Double() int {
	return int(i * 2)
}

func main() {
	myInt := Myint(10)
	fmt.Println(myInt.Double())
}
