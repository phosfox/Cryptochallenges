package main

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"log"
)

//SingleByteXor tries to decrypt the encoded string by xoring against a single byte
func SingleByteXor(enc string) (key int, message string) {
	decoded, err := hex.DecodeString(enc)
	if err != nil {
		log.Fatal(err)
	}
	xor := make([]byte, len(decoded))
	for i := 0; i < len(decoded); i++ {
		xor[i] = decoded[i] ^ 'b'
	}
	return 'a', hex.EncodeToString(xor)
}

func XorWithChar(enc string, key int) (dec string) {
	decoded, err := hex.DecodeString(enc)
	if err != nil {
		log.Fatal(err)
	}
	xor := make([]byte, len(decoded))
	for i := 0; i < len(decoded); i++ {
		xor[i] = decoded[i] ^ byte(key)
	}
	return hex.EncodeToString(xor)
}

//Xor takes two strings of the same length and xors them bitwise, returning the result
func Xor(a, b string) (string, error) {
	if len(a) != len(b) {
		return "", errors.New("strings have to be of the same length")
	}
	decodedA, err := hex.DecodeString(a)
	if err != nil {
		log.Fatal(err)
	}
	decodedB, err := hex.DecodeString(b)
	if err != nil {
		log.Fatal(err)
	}
	xor := make([]byte, len(decodedA))
	for i := 0; i < len(decodedA); i++ {
		xor[i] = decodedA[i] ^ decodedB[i]
	}
	return hex.EncodeToString(xor), nil
}

func HexToBase64String(s string) string {
	decodedSrc, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	encB64 := base64.StdEncoding.EncodeToString(decodedSrc)
	return encB64
}

func hexToBase64(encHex []byte) []byte {
	decHex := make([]byte, hex.DecodedLen(len(encHex)))
	_, err := hex.Decode(decHex, encHex)
	if err != nil {
		log.Fatal(err)
	}
	encBase64 := make([]byte, base64.StdEncoding.EncodedLen(len(decHex)))
	base64.StdEncoding.Encode(encBase64, decHex)
	return encBase64
}
