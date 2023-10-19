package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck which is a slice of strings.
type deck []string

// This works like a constructor
func newDeck() deck{
	cards := deck{}
	
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

/*This code defines a function called deal that takes a deck of cards (d) and a hand size 
(handSize) as input. It then returns two decks: the first deck consists of the first handSize cards from d, 
and the second deck consists of the remaining cards in d after the first handSize cards.*/
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	splitDeck := strings.Split(string(bs), ",")

	return deck(splitDeck)
}


// generador random de cartas
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) // creamos un source para crear aleatoriedad, con el dìa 
	//corriente en nanosegundos (representados en int64)
	rand.New(source) // generador de numeros aleatorios
	
	for i := range d {
		newPosition := rand.Intn(len(d) - 1) // generador de numeros aleatorios
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}