package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// equeal with using colon equal (:=)
	// card := newCard()
	// to change value use colon aonly :
	// card = "Five of Diamonds"
	// fmt.Println((card))

	//initialize new array
	cards := []string{"Ace of Diamonds" , newCard()}
	// append to join array to new value , but make new array
	cards = append(cards,"six of spades")

	for i,card := range cards {
		
		fmt.Println(i,card)
	}

}

func newCard() string  {
	
	return "Five of Diamonds"

}
