package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	newDeck := newDeck()

	if len(newDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(newDeck))
	}

	if newDeck[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be the Ace of Spades. Instead got %v", newDeck[0])
	}

	if newDeck[len(newDeck)-1] != "King of Diamonds" {
		t.Errorf("Expected last card to be King of Diamonds. Instead got %v", newDeck[len(newDeck)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	newDeck := newDeck()

	newDeck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
