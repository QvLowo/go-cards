package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"One", "Two", "Three", "Four"}

	// 不需要使用的index 用 _ 代替
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// receiver func
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// deck轉成string拼接,並用逗號分隔
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// ioutil.WriteFile() is deprecated (since Go v1.16) use os.WriteFile instead.
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename) //bs = byte slice
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1) // quit the program
	}

	s := strings.Split(string(bs), ",") // []string
	return deck(s)                      // type conversions
}

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)        // random
		d[i], d[newPosition] = d[newPosition], d[i] // swap
	}
}
