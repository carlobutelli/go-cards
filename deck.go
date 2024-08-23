package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type 'deck'
// which is a slice of strings
type deck []string

/* Adding some function to the newly created type 'deck' */
func newDeck() deck {
	// := tells go to initialize a variable and find out which type this variable is
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Hunchback", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

/* (d deck) is a receiver, every variable of type deck will have the function 'print' */
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// It's not returning any value
// To be used like static on the actual slice i.e. cards.shuffle()
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source) // create a new random

	for i := range d {
		newPosition := r.Intn(len(d) - 1)           // generate a random number btw 0 and len(d)-1 to get a random index
		d[i], d[newPosition] = d[newPosition], d[i] // swap: take what's inside i and put it into newPosition, then take what's into newPosition and put it into i
	}
}

// returning multiple values
// in this case it returns 2 values of type 'deck' from one deck
func (d deck) deal(handSize int) (deck, deck) {
	fmt.Println("getting hand...")
	// data[startIndexInlcuding:upToNotIncluding] ==> data[0:2] will include item with index 0 and 1 but not item with index 2
	return d[:handSize], d[handSize:]
}

func (d deck) remove(index int) deck {
	return d[index:]
}

/* Save deck type object to a file */
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

/* it returns a string from an object of type 'deck' through a joint of items in a slice of string */
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) toByte() []byte {
	return []byte(d.toString())
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	check(err)
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func check(err error) {
	if err != nil {
		// if something goes really wrong you have 2 options:
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - Log the error and entirely quit the program
		fmt.Println("Error:", err)
		panic(err)
	}
}

func (d deck) shuffleRevised() deck {
	fmt.Println("shuffling cards...")
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
	return d
}

/* Some Theory */
/*
  function signature: func <receiver> funcName(<args>) <return value(s)> {}
  - receiver is what can the function work on
  - args are the input parameters passed to the function
  - return value(s) is the returning value (it can be skipped meaning the func does not return any value)
*/
