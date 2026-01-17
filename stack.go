package main

// Stack represents a stack of cards.
type Stack struct {
	Cards []Card
}

// Push adds a card to the top of the stack.
func (s *Stack) Push(c Card) {
	s.Cards = append(s.Cards, c)
}

// Pop removes and returns the top card of the stack.
func (s *Stack) Pop() (Card, bool) {
	if len(s.Cards) == 0 {
		return Card{}, false
	}
	c := s.Cards[len(s.Cards)-1]
	s.Cards = s.Cards[:len(s.Cards)-1]
	return c, true
}

// Peek returns the top card of the stack without removing it.
func (s *Stack) Peek() (Card, bool) {
	if len(s.Cards) == 0 {
		return Card{}, false
	}
	return s.Cards[len(s.Cards)-1], true
}

// IsEmpty checks if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.Cards) == 0
}

// Clear removes all cards from the stack.
func (s *Stack) Clear() {
	s.Cards = nil
}
