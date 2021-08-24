package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// user_password := "username:password"
	user_password := "abcabcabc" // Each 3 characters -> 4 characters
	fmt.Println(user_password)
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(user_password)))
}
