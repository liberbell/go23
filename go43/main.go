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
	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)

	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)

	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fs := r2.FindString("/view/test")
	fmt.Println(fs)
}
