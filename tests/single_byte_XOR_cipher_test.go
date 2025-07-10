package tests

import (
	"cryptopals/utils"
	"testing"
)

// Challenge 1 Set 3: Single-byte XOR cipher
// This test checks if the SingleByteXORCipher function correctly decrypts a hex string that
// has been XORed with a single byte
func TestSingleByteXORCipher(t *testing.T) {
	plaintext := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	expectedMessage := "Cooking MC's like a pound of bacon"

	ciphertext, err := utils.SingleByteXORCipher(plaintext)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if ciphertext != expectedMessage {
		t.Errorf("SingleByteXORCipher failed.\nExpected: %s\nGot: %s", expectedMessage, ciphertext)
	}
}