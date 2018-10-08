package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	s, err := hex.DecodeString("68616c6c6f")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(s))
}
