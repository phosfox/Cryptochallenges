package main

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegexpDecrypt(t *testing.T) {
	a := "Cooking MC's like a pound of bacon"
	b := "i<jomn<al`k=jllhm8iahkmjlakljo<=j=<onjmil`iih<j?ljl:<oknhijkTS"
	c := "M↑NKIJ↑EHDO↓NHHLI∟MELOINHEOHNK↑↓N↓↑KJNIMHDMML↑N←HNH▲↑KOJLMNOpw"
	strings := []string{a, b, c}
	pattern := regexp.MustCompile("^[a-zA-Z0-9,\x60'\\.\\s-\\+]+$")

	for _, s := range strings {
		if pattern.MatchString(s) {
			t.Errorf("%s matched", s)
		}
	}
}

func TestGetLetterFrequencyScore(t *testing.T) {
	a := "Cooking MC's like a pound of bacon"
	score := GetLetterFrequencyScore(a)
	t.Errorf("Score: %d", score)
}
func TestXor(t *testing.T) {
	a := "1c0111001f010100061a024b53535009181c"
	b := "686974207468652062756c6c277320657965"
	correctResult := "746865206b696420646f6e277420706c6179"
	result, err := Xor(a, b)
	if err != nil {
		fmt.Println("strings are not of the same length")
	}
	if correctResult != result {
		t.Errorf("Expected: %s, got %s", correctResult, result)
	}
}

func TestXorWithChar(t *testing.T) {
	s := "aaaaaaa"
	key := 'b'
	message := XorWithChar([]byte(s), key)
	t.Errorf("Got: %s \n as String %s", message, string(message))
}
