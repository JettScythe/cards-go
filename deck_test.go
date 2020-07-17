package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := new()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 but got %v", len(d))
	}
	firstCard := Card{"Ace", "Spades"}
	if d[0] != firstCard {
		t.Errorf("Expected Ace Spades but got %v", d[0])
	}
	lastCard := Card{"King", "Clubs"}
	if d[len(d)-1] != lastCard {
		t.Errorf("Expected King Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToFileandNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := new()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
