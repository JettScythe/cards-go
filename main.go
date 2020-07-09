package main

func main() {
	// create a new deck
	cards := newDeck()
	// shuffle the deck
	cards.shuffle()
	// deal a hand of 5 cards
	hand, _ := deal(cards, 5)
	// print cards to console
	hand.print()
}
