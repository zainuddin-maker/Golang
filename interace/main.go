package main

import "fmt"

type bot interface{
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb:= spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())

}


// printGreeting will error red , becasue in golang we can make function with sama name , but can with different receiver.

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())

// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())

// }

// we can use just type without value , if the value not used.
func (englishBot) getGreeting() string {

	return "Hi There"
}

func (spanishBot) getGreeting() string {

	return "Hola"
}

