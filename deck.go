package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func (d deck) print() {
	for _, card := range d {

		fmt.Println(card)
	}
}

func newDeck() deck {
	d := deck{}
	cardSuits := []string{"Diamonds", "Sprades", "Hearts", "Clubs"}
	cardNums := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	for _, cardSuit := range cardSuits {
		for _, cardNum := range cardNums {
			d = append(d, cardNum+" of "+cardSuit)
		}
	}
	return d
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))

}

func deal(d deck, handSize int) (deck, deck) {
	if len(d) < handSize {
		return deck{}, d
	}
	return d[handSize:], d[:handSize]
}

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d))
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
