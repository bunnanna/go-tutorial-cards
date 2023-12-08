package main

import "fmt"

type deck []string

func (d deck) print() {
	for _, card := range d {

		fmt.Println( card)
	}
}

func newDeck() deck {
	d := deck{}
	cardSuits := []string{"Diamonds","Sprades","Hearts","Clubs"}
	cardNums := []string{"A","2","3","4","5","6","7","8","9","10","J","Q","K"}
	for _,cardSuit := range cardSuits{
		for _,cardNum := range cardNums{
			d = append(d, cardNum+" of "+cardSuit)
		}
	}
	return d
}

func deal(d deck, handSize int) (deck,deck) {
	if len(d)<handSize {
		return deck{},d
	}
	return d[handSize:],d[:handSize]
}