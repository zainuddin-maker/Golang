package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	
	cards := deck{}

	cardSuits := []string{"Spades","Diamonds","Hearts"}
	cardValues := []string{"Ace","Two","Three","Four"}

	for i,suit := range cardSuits {

		
		
	}
}

func (d deck) print() {
	for i, card := range d {

		fmt.Println(i, card)
	}
}


