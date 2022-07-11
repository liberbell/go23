package main

import "net/http"

func main() {
	resp, _ := http.Get("https://google.com")
	defer resp.Body.Close()
}
