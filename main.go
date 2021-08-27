package main

import (
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	// user_password := "username:password"
	user_password := "abcabcabc" // Each 3 characters -> 4 characters
	fmt.Println(user_password)
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(user_password)))
	// Now following https://golangdocs.com/aes-encryption-decryption-in-golang
	// cipher key
	key := "thisis32bitlongpassphraseimusing"

	// plaintext
	pt := "This is a secret!"

	c := EncryptAES([]byte(key), pt)
	fmt.Println("Encrypted:", c)
	// decrypt
	DecryptAES([]byte(key), c)
}

func EncryptAES(key []byte, plaintext string) string {
	// create cipher
	c, err := aes.NewCipher(key)
	CheckError(err)

	// allocate space for ciphered data
	out := make([]byte, len(plaintext))

	// encrypt
	c.Encrypt(out, []byte(plaintext))
	// return hex string
	return hex.EncodeToString(out)
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
}

func DecryptAES(key []byte, ct string) {
	ciphertext, _ := hex.DecodeString(ct)

	c, err := aes.NewCipher(key)
	CheckError(err)

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)

	s := string(pt[:])
	fmt.Println("DECRYPTED:", s)
}
