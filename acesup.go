package main

const STACKS = 4

// AcesUp represents the state of an Aces Up game.
type AcesUp struct {
	Deck        Deck
	Tableau     Tableau
	DiscardPile Stack
}

// NewGame initializes a new game of Aces Up.
func (game *AcesUp) Init() {
	game.Tableau.Clear()
	game.DiscardPile.Clear()
	game.Deck.Reset()
	game.Deck.Shuffle()
}

// Deal deals one card to each tableau stack.
func (game *AcesUp) Deal() (cardsDealt []Card, isDeckEmpty bool) {
	for i := 0; i < SUITS; i++ {
		if card, ok := game.Deck.Pop(); ok {
			game.Tableau.Stacks[i].Push(card)
			cardsDealt = append(cardsDealt, card)
		} else {
			isDeckEmpty = true
		}
	}
	return
}

// CanMove checks if the top card of fromStack can be moved to another empty stack.
func (game *AcesUp) CanMove(fromStack int) (bool, int) {

	// Validate fromStack index
	if fromStack < 0 || fromStack >= STACKS {
		return false, -1
	}

	// Check if fromStack is empty
	if game.Tableau.Stacks[fromStack].IsEmpty() {
		return false, -1
	}

	// Find an empty stack to move to
	for i := 0; i < STACKS; i++ {
		if i == fromStack {
			continue
		}
		if game.Tableau.Stacks[i].IsEmpty() {
			return true, i
		}
	}

	return false, -1
}

// CanDiscard checks if the top card of the specified stack can be discarded.
func (game *AcesUp) CanDiscard(stackIndex int) (bool, Card) {
	topCards := game.Tableau.Peek()

	// Validate stack index
	if stackIndex < 0 || stackIndex >= STACKS {
		return false, Card{}
	}
	if game.Tableau.Stacks[stackIndex].IsEmpty() {
		return false, Card{}
	}

	card := topCards[stackIndex]

	// Cannot discard an Ace
	if card.Rank == ACE {
		return false, card
	}

	// Check for higher card of same suit in other stacks
	for i := 0; i < STACKS; i++ {
		if i == stackIndex || game.Tableau.Stacks[i].IsEmpty() {
			continue
		}
		other := topCards[i]
		if other.Suit == card.Suit && other.Rank > card.Rank {
			return true, card
		}
	}
	return false, card
}

// Discard removes the top card from the specified stack and places it in the discard pile.
func (game *AcesUp) Discard(stackIndex int) (bool, Card) {
	if canDiscard, card := game.CanDiscard(stackIndex); !canDiscard {
		return false, card
	}

	card, _ := game.Tableau.Stacks[stackIndex].Pop()
	game.DiscardPile.Push(card)
	return true, card
}

// Move transfers the top card from fromStack to the next available empty stack.
func (game *AcesUp) Move(fromStack int) bool {

	// Check if move is possible and get target stack index
	canMove, toStack := game.CanMove(fromStack)
	if !canMove {
		return false
	}

	// Perform the move
	card, _ := game.Tableau.Stacks[fromStack].Pop()
	game.Tableau.Stacks[toStack].Push(card)
	return true
}

// CheckWinLoose checks for win or loose conditions.
func (game *AcesUp) CheckWinLoose() (win bool, loose bool) {

	// Check for win condition: all four aces are on the tableau
	aceCount := 0
	for i := 0; i < STACKS; i++ {
		if game.Tableau.Stacks[i].IsEmpty() {
			continue
		}
		topCard, _ := game.Tableau.Stacks[i].Peek()
		if topCard.Rank == ACE {
			aceCount++
		}
	}
	win = (aceCount == SUITS)

	// Check for loose condition: deck is empty
	loose = game.Deck.IsEmpty()

	return
}
