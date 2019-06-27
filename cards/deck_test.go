package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %d", len(d))
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected first element not equal to %s", d[0])
	}

	if d[len(d)-1] != "Four of Hearts" {
		t.Errorf("Expected first element not equal to %s", d[len(d)-1])
	}
}

func TestSavetofilenewdeckfromfile(t *testing.T) {
	d := newDeck()
	os.Remove("_decktesting")
	d.saveToFile("_decktesting")
	loadeddeck := newDeckFromFile("_decktesting")
	if len(loadeddeck) != 16 {
		t.Errorf("Expected deck length of 16 but got %d", len(loadeddeck))
	}
	os.Remove("_decktesting")
}
