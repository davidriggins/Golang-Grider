package main

import "fmt"

// Create a new type of 'deck' which is a slice of strings
// It is like the new "deck" type kind of extends the string
// behavior
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}