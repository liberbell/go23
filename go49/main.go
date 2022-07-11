package main

type Person struct {
	Name      string
	Age       int
	nicknames []string
}

func main() {
	b := []byte(`{"name":"Elton", "Age":72, "nicknames": ["a", "b", "c"]}`)

}
