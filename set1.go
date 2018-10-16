package main

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

//GetLetterFrequencyScore calculates the letter frequency of a given string
func GetLetterFrequencyScore(s string) int {
	score := 0
	s = strings.ToLower(s)
	for _, c := range s {
		switch c {
		case 'e':
			score += 12
		case 'a':
			score += 8
		case 'o':
			score += 7
		case 'i':
			score += 7
		case 'h':
			score += 6
		case 't':
			score += 9
		case 's':
			score += 6
		case 'r':
			score += 6
		}
	}
	return score
}

//RegexpDecryptWithSingleByteXorCipher tries to decrypt the encoded string by xoring against a single byte
func RegexpDecryptWithSingleByteXorCipher(enc []byte) map[int][]byte {
	pattern := regexp.MustCompile("^[a-zA-Z0-9,\x60'\\.\\s-\\+]+$")
	m := make(map[int][]byte)
	//32 first printable char 126 last
	for i := 65; i <= 122; i++ {
		xor := XorWithChar(enc, rune(i))
		if pattern.MatchString(string(xor)) {
			m[i] = xor
		}
	}
	return m
}

//DecryptWithSingleByteXorCipher tries to decrypt the encoded string by xoring against a single byte
func DecryptWithSingleByteXorCipher(enc []byte) map[int][]byte {
	m := make(map[int][]byte)
	//32 first printable char 126 last
	for i := 32; i <= 126; i++ {
		xor := XorWithChar(enc, rune(i))
		m[i] = xor
	}
	return m
}

//XorWithChar xors a given string with a rune and returns that
func XorWithChar(enc []byte, key rune) []byte {
	xor := make([]byte, len(enc))
	for i := 0; i < len(enc); i++ {
		xor[i] = enc[i] ^ byte(key)
	}
	return xor
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

func hexToBytes(encHex []byte) []byte {
	decHex := make([]byte, hex.DecodedLen(len(encHex)))
	i, err := hex.Decode(decHex, encHex)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}
	return decHex[:i]
}

func challenge4() {
	file, err := os.Open("challenge4.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	newFile, err := os.Create("dec.txt")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = line[:len(line)-2]
		hexLine := hexToBytes([]byte(line))
		m := DecryptWithSingleByteXorCipher(hexLine[:len(hexLine)-1])

		for key, val := range m {
			valAsString := string(val)
			score := GetLetterFrequencyScore(valAsString)
			if score > 100 {
				s := "String: " + valAsString + " Key: " + string(key) + "\n"
				newFile.WriteString(s)
			}
		}
	}

}
