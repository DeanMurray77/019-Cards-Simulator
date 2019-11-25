package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	var hand deck

	hand, cards = deal(cards, 5)
	hand.print()
	fmt.Println("--")
	cards.print()

	cards.saveToFile("cards.dat")

}
