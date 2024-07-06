package main

import (
	"os"
	"testing"
)

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

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	fileName := "_decktesting"
	os.Remove(fileName)

	deck := newDeck()
	deck.saveToFile(fileName)
	loadedDeck := newDeckFromFile(fileName)
	os.Remove(fileName) // cleans up test files, comment to view generated file

	length := len(loadedDeck)
	if length != 16 {
		t.Errorf("Expected deck length of 16, but got %v", length)
	}
}
