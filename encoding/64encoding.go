package main

import "fmt"
import b64 "encoding/base64"
import b32 "encoding/base32"

func main() {
	data := "1022094818"
	enc64 := b64.StdEncoding.EncodeToString([]byte(data))
	enc32 := b32.StdEncoding.EncodeToString([]byte(data))

	fmt.Println("64 encoding: ", enc64)
	fmt.Println("32 encoding: ", enc32)
}
