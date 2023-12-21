package main

func main() {
	// cards := newDeck()
	cards := newDeckFromFile("first_deck.txt")

	// err := cards.saveToFile("first_deck.txt")
	// if err != nil {
	// 	return
	// }
	// cards, hand := deal(cards, 5)
	// hand.print()
	cards.shuffle()
	cards.print()

}
