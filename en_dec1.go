package main

import (
    "encoding/base64"
    "fmt"
)

func main() {

    // String to encode
    str := "Hello World!"

    // base64.StdEncoding: Standard encoding with padding
    // It requires a byte slice so we cast the string to []byte
    encodedStr := base64.StdEncoding.EncodeToString([]byte(str))
    fmt.Println("Encoded:", encodedStr)

    // Decoding may return an error, in case the input is not well formed
    decodedStr, err := base64.StdEncoding.DecodeString(encodedStr)
    if err != nil {
        panic("malformed input")
    }
    fmt.Println("Decoded:", string(decodedStr))
}