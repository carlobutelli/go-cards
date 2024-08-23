package main

func main() {

	cards := newDeck()

	// cards.print()

	cards.shuffle()
	// cards.shuffleRevised()

	// cards.print()

	// cards.deal(4)

	// hand.print()
	// remainingCards.print()

	// str := "Hi there!"
	// fmt.Println([]byte(str))

	// cards.saveToFile("myDeck")

	// cards := newDeckFromFile("myDeck1")
	cards.print()

}
