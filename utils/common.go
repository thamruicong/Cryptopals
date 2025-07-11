package utils

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func HexToBase64(hexStr string) (string, error) {
	// Decode the hex string
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", fmt.Errorf("Error decoding hex: %v", err)
	}

	// Encode the bytes to base64
	base64Str := base64.StdEncoding.EncodeToString(bytes)

	return base64Str, nil
}

func FixedXOR(a, b string) (string, error) {
	aBytes, err := hex.DecodeString(a)
	if err != nil {
		return "", fmt.Errorf("Error decoding hex %q: %v", a, err)
	}

	bBytes, err := hex.DecodeString(b)
	if err != nil {
		return "", fmt.Errorf("Error decoding hex %q: %v", b, err)
	}

	if len(aBytes) != len(bBytes) {
		return "", fmt.Errorf("Input strings must be of equal length: %d vs %d", len(aBytes), len(bBytes))
	}

	result := make([]byte, len(aBytes))
	for i := range result {
		result[i] = aBytes[i] ^ bBytes[i]
	}

	return hex.EncodeToString(result), nil
}

func RepeatingKeyXOR(plaintext, key string) (string, error) {
	if len(key) == 0 {
		return "", fmt.Errorf("Key must not be empty")
	}

	ciphertext := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i++ {
		ciphertext[i] = plaintext[i] ^ key[i % len(key)]
	}

	return hex.EncodeToString(ciphertext), nil
}

func HammingDistance(a, b string) (int, error){
	if len(a) != len(b) {
		return 0, fmt.Errorf("Strings must be of equal length: %d vs %d", len(a), len(b))
	}

	var distance int
	for i := range a {
		xor := a[i] ^ b[i]
		for j := 0; j < 8; j++ {
			if xor & (1 << j) != 0 {
				distance++
			}
		}
	}
	
	return distance, nil
}