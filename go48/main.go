package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	// resp, _ := http.Get("https://google.com")
	// defer resp.Body.Close()

	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	base, err := url.Parse("https://google.com/dfoujlsadf")
	if err != nil {
		fmt.Println(err)
	}
	reference, _ := url.Parse("/test?a=1&b=2")
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)

	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("Any-header", "apple/chrome")
	q := req.URL.Query()
	fmt.Println(q)
}
