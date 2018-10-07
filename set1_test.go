package main

import (
	"fmt"
	"testing"
)

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
