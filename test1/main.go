package main

import (
	"fmt"
	"strconv"
)

func main() {
	f := 1.11
	conv_f, _ := strconv.Atoi(f)
	fmt.Println(conv_f)
}
