package utils

import (
	"encoding/hex"
	"fmt"
	"unicode"
)

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