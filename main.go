package main

import (
	"encoding/hex"
	"fmt"
)

type ding struct {
	encKey int
	msg    string
	score  int
}

func main() {
	clrtxt := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := "ICE"
	encrTxt := XorWithString(clrtxt, key)
	fmt.Println(encrTxt)
	hexEcrTxt := hex.EncodeToString(encrTxt)
	fmt.Println(hexEcrTxt)
}

//XorWithString implements a repeating-key XOR
func XorWithString(msg, key string) []byte {
	lenMsg := len(msg)
	lenKey := len(key)
	encMsg := make([]byte, lenMsg)

	for i := 0; i < lenMsg; i++ {
		encMsg[i] = msg[i] ^ key[i%lenKey]
		//fmt.Printf("%c = %c ^ %c\n", encMsg[i], msg[i], key[i%lenKey])
	}
	return encMsg
}
