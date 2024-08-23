package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) { // t is our test handler
	// create new Deck
	d := newDeck()
	// test the newly created Deck has correct number of cards
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 but got %v", len(d))
	}
	// test
}

func TestFirstCard(t *testing.T) {
	d := newDeck()
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades, but got %v", d[0])
	}
}

func TestLastCard(t *testing.T) {
	d := newDeck()
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loaded := newDeckFromFile("_decktesting")

	if len(loaded) != 52 {
		t.Errorf("Expected 52 cards in deck, but got %v", len(loaded))
	}
	os.Remove("_decktesting")

}
