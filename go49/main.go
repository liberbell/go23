package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string
	Age       int
	Nicknames []string
}

func main() {
	b := []byte(`{"name":"Elton", "age":72, "nicknames": ["a", "b", "c"]}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)

}