package main

func main() {
	// cards := newDeck()
	cards := readFromFile("first_deck.txt")

	// cards.saveToFile("first_deck.txt")
	// cards, hand := deal(cards, 5)
	// hand.print()
	cards.print()

}
