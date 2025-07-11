package tests

import (
	"cryptopals/utils"
	"os"
	"strings"
	"testing"
)

// Set 1 Challenge 4: Detect single-character XOR
// This test checks if the SingleByteXORCipher function can correctly identify the line
// with the highest score in a specific file containing hex-encoded strings
func TestDetectSingleCharXOR(t *testing.T) {
	data, err := os.ReadFile("../../assets/set1_challenge4.txt")
	if err != nil {
		t.Fatalf("Unable to read file: %v", err)
	}

	lines := strings.Split(string(data), "\n")

	var bestLine string
	var bestScore float64
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		decrypted, score, err := utils.SingleByteXORCipher(line)
		if err != nil {
			t.Errorf("Error decrypting line %q: %v", line, err)
			continue
		}
		if score > bestScore {
			bestScore = score
			bestLine = decrypted
		}
	}
	if bestLine == "" {
		t.Error("No valid lines found or all lines failed decryption")
		return
	}
	expectedMessage := "Now that the party is jumping\n"
	if bestLine != expectedMessage {
		t.Errorf("Expected: %s\nGot: %s", expectedMessage, bestLine)
	}
}