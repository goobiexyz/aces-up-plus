package main

import (
	"math/rand"
	// "fmt"
)

type Deck struct {
	Stack
}

// Creates a standard 52-card deck.
func (d *Deck) Reset() {
	d.Cards = nil

	for suit := 1; suit <= 4; suit++ {
		for rank := 2; rank <= 13; rank++ {
			d.Cards = append(d.Cards, Card{Rank: rank, Suit: suit})
		}
	}
}

// Shuffles the deck
func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}
