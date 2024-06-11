package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	// This is generally taken from gobyexample but with a few other things

	s := "this is a string I will hash"

	// Initialize the hash. Note that you can pretty much use any hash
	// that Go supports here, like MD5, SHA512, etc.
	h := sha256.New()

	// Hashing requires a byte string. Good thing you can do that with
	// Go fairly easily
	h.Write([]byte(s))

	// The hash output is stored under h.Sum(nil), where you can append
	// something to the hash here if you'd like. Not really all that necessary.
	bs := h.Sum(nil)

	// If we print the `bs` variable it will be gibberish and unprintable
	// characters. The generally accepted way to print a hash is by printing
	// the hex format

	var hexString string = hex.EncodeToString(bs)
	fmt.Println("Hex representation:", hexString)

	// You can also b64 encode it if you'd like.
	var hexB64 string = base64.StdEncoding.EncodeToString(bs)
	fmt.Println("Base64:", hexB64)

	// Otherwise you can Base64 URL encode as well, which should get
	// rid of some of the weirdo characters that don't print well in
	// a browser's GET parameters
	var hexUB64 string = base64.URLEncoding.EncodeToString(bs)
	fmt.Println("Base64 URL safe:", hexUB64)

	// Similarly, you can decode it as well
	decoded, err := base64.URLEncoding.DecodeString(hexUB64)
	if err != nil {
		panic(err)
	}
	fmt.Println("This won't print well but it's decoded:", string(decoded))
}
