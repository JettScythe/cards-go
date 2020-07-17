package main

func main() {
	// Create a new deck
	newDeck := new()
	// Shuffle the deck
	newDeck.shuffle()
	// Deal a hand of 5 cards
	hand, _ := deal(newDeck, 5)
	// print hand of cards
	hand.print()
	// save hand of cards
	hand.saveToFile("User1")
	// read hand from file
	savedHand := newDeckFromFile("User1")
	savedHand.print()
}
