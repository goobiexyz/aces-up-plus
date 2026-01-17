package main

import "fmt"

var game AcesUp

func main() {

	game.Init()
	game.Deal()
	//printTableau()

	won, lost := false, false

	// Main game loop
	for {
		printTableau()
		fmt.Print("> ")

		// Read user command
		var cmd string
		if _, err := fmt.Scan(&cmd); err != nil {
			fmt.Println("Input error, exiting.")
			return
		}

		// Process command
		switch cmd {

		// draw, deal
		case "draw", "deal":
			hand, isDeckEmpty := game.Deal()

			// Report dealt cards
			if len(hand) > 0 {
				fmt.Print("Dealt " + fmt.Sprint(len(hand)) + " card")
				if len(hand) > 1 {
					fmt.Print("s")
				}
				fmt.Println(".")
			}

			// Report if deck is empty
			if isDeckEmpty {
				fmt.Println("The deck is empty.")
			}

		// discard
		case "discard", "disc":
			var idx int

			// Validate index input
			if _, err := fmt.Scan(&idx); err != nil {
				fmt.Println("Invalid index.")
				continue
			}
			if idx < 1 || idx > STACKS {
				fmt.Println("Invalid index.")
				continue
			}

			// Perform the discard
			success, card := game.Discard(idx - 1)
			if success {
				fmt.Println("Discarded.")
			} else {
				fmt.Println("You cannot discard that " + card.String() + ".")
			}

		case "move":
			var idx int

			// Validate index input
			if _, err := fmt.Scan(&idx); err != nil {
				fmt.Println("Invalid index.")
				continue
			}
			if idx < 1 || idx > STACKS {
				fmt.Println("Invalid index.")
				continue
			}

			// Performs the move.
			success := game.Move(idx - 1)
			if success {
				fmt.Println("Moved.")
			} else {
				fmt.Println("Nowhere to move that card.")
			}
		default:
			fmt.Println("Unknown command.")
		}

		won, lost = game.CheckWinLoose()
		if won || lost {
			break
		}
	}

	if won {
		fmt.Println("You won!")
	} else if lost {
		fmt.Println("You lost.")
	}
}

func printTableau() {
	for i := 0; i < STACKS && i < len(game.Deck.Cards); i++ {
		c, cardExists := game.Tableau.Stacks[i].Peek()
		if i > 0 {
			fmt.Print(" ")
		}

		if !cardExists {
			fmt.Print("[  ]")
			continue
		}

		if len(game.Tableau.Stacks[i].Cards) > 1 {
			fmt.Print("[")
		}

		fmt.Printf(c.String())
	}
	fmt.Println()
}
