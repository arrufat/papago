// +build ignore

// decode papago http payload from base64

package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a base64 string.")
		return
	}
	str := os.Args[1]
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%q\n", data)
}
