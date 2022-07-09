package main

import (
	"fmt"
	"regexp"
)

func main() {
	// t := time.Now()
	// fmt.Println(t)
	// fmt.Println(t.Format(time.RFC3339))
	// fmt.Println(t.Year())
	// fmt.Println(t.Month())
	// fmt.Println(t.Hour())
	match, _ := regexp.MatchString("a([a-z0-9]+)e", "appl0e")
	fmt.Println(match)
}
