package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

var DB = map[string]string{
	"User1Key": "User1Secret",
	"User2Key": "User2Secret",
}

func Server(apikey, sign string, data []byte) {
	apiSecret := DB[apikey]
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	expectedHMAC := hex.EncodeToString(h.Sum(nil))
}

func main() {
	const apiKey = "User1Key"
	const apiSecret = "User2Secret"
	data := []byte("data")

	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	// fmt.Println(h)
	sign := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sign)

	Server(apiKey, sign, data)
}
