package main

import "fmt"

func main() {
	hand := newDeck()
	hand.print()
	fmt.Println("-----")
	hand.shuffle()
	hand.print()
}
