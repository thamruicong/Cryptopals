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