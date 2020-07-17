package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

type Card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
}
type Deck []Card

// Create a new deck
func new() (deck Deck) {
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	for i := range values {
		for j := range suits {
			card := Card{
				Value: values[i],
				Suit:  suits[j],
			}
			deck = append(deck, card)
		}
	}
	return
}

func (d Deck) print() {
	for _, card := range d {
		fmt.Printf("%v of %v\n", card.Value, card.Suit)
	}
}

// Deal a specified amount of cards
func deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

// Shuffle the deck
func (d Deck) shuffle() Deck {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
	return d
}

// Save the deck to a file
func (d Deck) saveToFile(filename string) error {
	var jsonData0 []byte
	jsonData0, err0 := json.Marshal(d)
	if err0 != nil {
		log.Println(err0)
	}
	return ioutil.WriteFile(filename, jsonData0, 0666)
}

// Read saved hand from file, if there is no saved hand, call a new deck, shuffle deal, and use that
func newDeckFromFile(filename string) Deck {
	bs, err1 := ioutil.ReadFile(filename)
	if err1 != nil {
		fmt.Println("Error:", err1, "\nCreating new deck")
		createDeck := new()
		createDeck.shuffle()
		createHand, _ := deal(createDeck, 5)
		return createHand
	}
	var savedDeck Deck
	err2 := json.Unmarshal(bs, &savedDeck)
	if err2 != nil {
		fmt.Println(err2)
	}
	return savedDeck
}
