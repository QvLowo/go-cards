package main

import "fmt"

type deck []string

// receiver func
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
