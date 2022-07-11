package main

import (
	"fmt"
	"net/url"
)

func main() {
	// resp, _ := http.Get("https://google.com")
	// defer resp.Body.Close()

	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	base, err := url.Parse("https://goog le.com")
	if err != nil {
		fmt.Println(err)
	}
	reference, _ := url.Parse("/")
	fmt.Println(base)
}
