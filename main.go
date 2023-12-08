package main

func main() {
	cards := newDeck()
	cards, hand := deal(cards, 5)
	hand.print()
	cards.print()


}
