package main

func main() {
	cards := newDeck()
	cards.shuffle()
	myhand, remaining := deal(cards, 7)
	myhand.saveDeck("my")
	remaining.saveDeck("rem")

	my := loadDeck("my")
	// rem := loadDeck("rem")
	my.print()
	// rem.print()
}
