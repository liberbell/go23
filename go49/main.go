package main

import (
	"encoding/json"
	"fmt"
)

type T struct {
}

type Person struct {
	Name      string   `json:"name,omitempty"`
	Age       int      `json:"age,omitempty"`
	Nicknames []string `json:"nicknames"`
	T         T        `json:"T"`
}

func main() {
	b := []byte(`{"name":"Elton", "age":20, "nicknames": ["a", "b", "c"]}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)

	v, _ := json.Marshal(p)
	fmt.Println(string(v))
}
