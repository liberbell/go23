package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func main() {
	const apiKey = "User1Key"
	const apiSecret = "User1Secret"

	h := hmac.New(sha256.New, []byte(apiSecret))
	fmt.Println(h)
}
