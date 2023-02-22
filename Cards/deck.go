package main

import "fmt"

// Create a new tyoe of 'deck' which is a slice of strings
type deck []string

func (d deck) newDeck() deck{
	return []
}


func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

