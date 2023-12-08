package main

func main() {
	cards := newDeck()
	cards = append(cards, "6 ha")
	cards.print()
}
