package tests

import (
	"cryptopals/utils"
	"testing"
)

// Challenge 1 Set 2: Fixed XOR
// This test checks if the FixedXOR function correctly computes the XOR of two hex strings
func TestFixedXOR(t *testing.T) {
	hexStrA := "1c0111001f010100061a024b53535009181c"
	hexStrB := "686974207468652062756c6c277320657965"
	expectedXOR := "746865206b696420646f6e277420706c6179"

	result, err := utils.FixedXOR(hexStrA, hexStrB)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != expectedXOR {
		t.Errorf("FixedXOR failed.\nExpected: %s\nGot: %s", expectedXOR, result)
	}
}