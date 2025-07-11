package tests

import (
	"cryptopals/utils"
	"testing"
)

// Set 1 Challenge 5: Implement repeating-key XOR
// This test checks if the RepeatingKeyXOR function can correctly encrypt a plaintext with a repeating key
func TestRepeatingKeyXOR(t *testing.T) {
	plaintext := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := "ICE"
	expectedCiphertext := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	ciphertext, err := utils.RepeatingKeyXOR(plaintext, key)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if ciphertext != expectedCiphertext {
		t.Errorf("RepeatingKeyXOR failed.\nExpected: %s\nGot: %s", expectedCiphertext, ciphertext)
	}
}