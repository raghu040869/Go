package main

import "fmt"

func main() {

	cards := newDeck()

	cards.print()
	fmt.Println("---------")
	hand, remainingCards := deal(cards, 5)
	hand.print()
	fmt.Println("---------")
	remainingCards.print()
	fmt.Println("---------")
	print(cards.toString())
	cards.saveToFile("ragstesting")
	fmt.Println("new cards---------")
	newcards := newDeckFromFile("ragstesting")
	newcards.print()
	fmt.Println(" shuffle---------")
	cards.shuffle()
	cards.print()
	fmt.Println("---------")
}
