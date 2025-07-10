package tests

import (
	"cryptopals/utils"
	"testing"
)

// Challenge 1 Set 1: Convert hex to base64
// This test checks if the HexToBase64 function correctly converts a hex string to a base
func TestHexToBase64(t *testing.T) {
	hexStr := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expectedBase64 := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	result, err := utils.HexToBase64(hexStr)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != expectedBase64 {
		t.Errorf("HexToBase64 failed.\nExpected: %s\nGot: %s", expectedBase64, result)
	}
}