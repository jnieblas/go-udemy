package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, fmt.Sprintf("%s of %s", value, suit))
		}
	}

	return cards
}

func newDeckFromFile(fileName string) deck {
	bs, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return strings.Split(string(bs), ",")
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		randI := r.Intn(len(d))
		d[i], d[randI] = d[randI], d[i]
	}
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}
