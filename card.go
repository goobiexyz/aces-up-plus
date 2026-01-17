package main

import "fmt"

const JACK = 10
const QUEEN = 11
const KING = 12
const ACE = 13
const HEARTS = 1
const DIAMONDS = 2
const CLUBS = 3
const SPADES = 4

const SUITS = 4
const RANKS = 12

type Card struct {
	Rank int
	Suit int
}

// Returns the rank of the card as a string. Option to abbreviate face cards.
func (c *Card) GetRankString(abbrev bool) string {
	if c.Rank == 0 {
		return "unset rank"

	} else if c.Rank < 10 && c.Rank > 1 {
		return fmt.Sprint(c.Rank)

	} else {
		switch c.Rank {
		case JACK:
			if abbrev {
				return "J"
			} else {
				return "Jack"
			}
		case QUEEN:
			if abbrev {
				return "Q"
			} else {
				return "Queen"
			}
		case KING:
			if abbrev {
				return "K"
			} else {
				return "King"
			}
		case ACE:
			if abbrev {
				return "A"
			} else {
				return "Ace"
			}
		default:
			return "invalid rank: " + fmt.Sprint(c.Rank)
		}
	}
}

// Returns the suit of the card as a string. Option to abbreviate face cards.
func (c *Card) GetSuitString(abbrev bool) string {
	switch c.Suit {
	case 0:
		return "unset suit"
	case HEARTS:
		if abbrev {
			return "H"
		} else {
			return "Hearts"
		}
	case DIAMONDS:
		if abbrev {
			return "D"
		} else {
			return "Diamonds"
		}
	case CLUBS:
		if abbrev {
			return "C"
		} else {
			return "Clubs"
		}
	case SPADES:
		if abbrev {
			return "S"
		} else {
			return "Spades"
		}
	default:
		return "invalid suit: " + fmt.Sprint(c.Suit)
	}
}

// Returns the Unicode symbol for the card's suit.
func (c *Card) GetSuitSymbol() string {
	switch c.Suit {
	case HEARTS:
		return "♥"
	case DIAMONDS:
		return "♦"
	case CLUBS:
		return "♣"
	case SPADES:
		return "♠"
	default:
		return "invalid suit: " + fmt.Sprint(c.Suit)
	}
}

func (c *Card) String() string {
	return fmt.Sprintf("[%s %s]", c.GetRankString(true), c.GetSuitSymbol())
}
