package main

import "testing"

func TestNewDeck(t *testing.T) {
	newDeck := newDeck()

	if len(newDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(newDeck))
	}
}
