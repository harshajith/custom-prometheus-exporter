package main

import (
	"example/cards/deck"
	"fmt"
)

func main() {

	fmt.Println("hello world")
	cards := deck.NewDeck()
	cards.Print()
	cards.SaveToFile("mydeck")
}

// func newCard() string {
// 	return fmt.Println("five of diamonds")
// }
