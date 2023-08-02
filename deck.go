package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

type deck []string

// Create a new deck
func newDeck() deck {
	cards := deck{}
	shape := []string{"Spade", "Clubs", "Heart", "Diamond"}
	number := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, s := range shape {
		for _, n := range number {
			cards = append(cards, s+" "+n)
		}
	}
	return cards
}

// Draws 'n' number of cards from the deck
func deal(d deck, n int) (deck, deck) {
	return d[:n], d[n:]
}

// Prints Deck
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// Converts deck to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Converts string to deck
func toDeck(s string) []string {
	return strings.Split(s, ",")
}

// Saves deck
func (d deck) saveDeck(filename string) {
	err := os.WriteFile(filename, []byte(d.toString()), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

// Loads deck
func loadDeck(filename string) deck {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return toDeck(string(data))
}

// Shuffles the deck
func (d deck) shuffle() deck {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
	return d
}
