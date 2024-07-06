package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	length := len(d)

	if length != 16 {
		t.Errorf("Expected deck length of 16, but got %v", length)
	}

	expectedFirstCard := "Ace of Spades"
	actualFirstCard := d[0]
	if actualFirstCard != expectedFirstCard {
		t.Errorf("Expected first card of %v, but got %v", expectedFirstCard, actualFirstCard)
	}

	expectedLastCard := "Four of Clubs"
	actualLastCard := d[length-1]
	if actualLastCard != expectedLastCard {
		t.Errorf("Expected last card of %v, but got %v", expectedLastCard, actualLastCard)
	}
}
