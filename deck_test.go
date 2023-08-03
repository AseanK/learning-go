package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Spade 1" || d[len(d)-1] != "Diamond K" {
		t.Errorf("Expected first card: Spade 1\n Got %v \n Expected last card: Diamond K \n Got %v", d[0], d[len(d)-1])
	}
}

func TestSaveDeckAndLoadDeck(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveDeck("_decktesting")
	ld := loadDeck("_decktesting")

	if len(ld) != 52 {
		t.Errorf("Expected deck length: 52, but got %v", len(ld))
	}

	os.Remove("_decktesting")
}

func TestShuffle(t *testing.T) {
	d1 := newDeck()
	d2 := newDeck()

	d1.shuffle()
	d2.shuffle()

	if d1[0] == d2[0] && d1[len(d1)-1] == d2[len(d2)-1] {
		t.Errorf("Expected different values")
	}
}
