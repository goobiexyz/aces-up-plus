package main

// Tableau represents the tableau area with four stacks.
type Tableau struct {
	Stacks [STACKS]Stack
}

// Clear removes all cards from all tableau stacks.
func (t *Tableau) Clear() {
	for i := 0; i < STACKS; i++ {
		t.Stacks[i].Clear()
	}
}

// Peek returns the top cards of all tableau stacks.
func (t *Tableau) Peek() []Card {
	var tops []Card
	for i := 0; i < STACKS; i++ {
		if card, ok := game.Tableau.Stacks[i].Peek(); ok {
			tops = append(tops, card)
		} else {
			tops = append(tops, Card{})
		}
	}
	return tops
}
