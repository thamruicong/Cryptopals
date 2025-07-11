package tests

import (
	"cryptopals/utils"
	"testing"
)

// Set 1 Challenge 6: Break repeating-key XOR
func Test(t *testing.T) {
}

// This test checks if the HammingDistance function correctly calculates the distance
func TestHammingDistance(t *testing.T) {
	stringA := "this is a test"
	stringB := "wokka wokka!!!"
	expectedDist := 37

	dist, err := utils.HammingDistance(stringA, stringB)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if dist != expectedDist {
		t.Errorf("HammingDistance failed.\nExpected: %d\nGot: %d", expectedDist, dist)
	}
}