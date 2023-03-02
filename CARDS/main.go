package main

// import "fmt"

// type deck []string

func main() {
	// var card string = "Ace of Spades"
	// equeal with using colon equal (:=)
	// card := newCard()
	// to change value use colon aonly :
	// card = "Five of Diamonds"
	// fmt.Println((card))

	//initialize new array

	// cards :=  newDeck()
	// // append to join array to new value , but make new array
	// // cards = append(cards, "six of spades")

	// // cards.print()
	//   hand, remainingCards := deal(cards,5)

	//   hand.print()
	//   remainingCards.print()

	// for i, card := range cards {

	// 	fmt.Println(i, card)
	// }

	// greeting := "Hi there!"

	// fmt.Println([]byte(greeting))

	cards := newDeck()
	cards.saveToFile("my_cards")
	// fmt.Println(cards.toString())

}

// func newCard() string {

// 	return "Five of Diamonds"

// }
