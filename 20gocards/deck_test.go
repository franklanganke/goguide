package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck lenght of 16, but got %v", len(d))
	}

	firstCard := "Ace of Spades"
	if d[0] != firstCard {
		t.Errorf("Expected first card of %v, but got %v", firstCard, d[0])
	}

	lastcard := "Jack of Club"
	if d[len(d)-1] != lastcard {
		t.Errorf("Expected last card of %v, but got %v", lastcard, d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	filename := "_decktesting.csv"
	os.Remove(filename)

	deck := newDeck()
	deck.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove(filename)
}
