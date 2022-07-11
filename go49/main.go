package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string
	Age       int
	nicknames []string
}

func main() {
	b := []byte(`{"name":"Elton", "Age":72, "nicknames": ["a", "b", "c"]}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}

}
