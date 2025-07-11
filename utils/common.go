package utils

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"unicode"
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

func SingleByteXORCipher(input string) (string, float64, error) {
	inputBytes, err := hex.DecodeString(input)
	if err != nil {
		return "", 0.0, fmt.Errorf("Error decoding hex %q: %v", input, err)
	}

	var highestScore float64
	var bestDecrypted []byte

	for key := 0; key < 256; key++ {
		decrypted := make([]byte, len(inputBytes))
		for i := range inputBytes {
			decrypted[i] = inputBytes[i] ^ byte(key)
		}
		score := scoreText(decrypted)

		if score > highestScore {
			highestScore = score
			bestDecrypted = decrypted
		}
	}
	return string(bestDecrypted), highestScore, nil
}

func scoreText(text []byte) float64 {
	frequency := map[byte] float64 {
		'e': 12.7, 't': 9.1, 'a': 8.2, 'o': 7.5, 'i': 7.0,
		'n': 6.7, ' ': 13.0, 's': 6.3, 'h': 6.1, 'r': 6.0,
		'd': 4.2, 'l': 4.0, 'u': 2.8, 'c': 2.8, 'm': 2.4,
	}
	var total float64
	for _, c := range text {
		total += frequency[byte(unicode.ToLower((rune(c))))]
	}
	return total
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