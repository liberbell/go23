package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	const apiKey = "User1Key"
	const apiSecret = "User1Secret"

	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write([]byte("data"))
	// fmt.Println(h)
	sign := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sign)
}
