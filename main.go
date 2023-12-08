package main

func main() {
	cards := deck{"A sp", newCard()}
	cards = append(cards, "6 ha")
	cards.print()
}

func newCard() string {
	return "5 di"
}