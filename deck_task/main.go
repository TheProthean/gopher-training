package main

import (
	"fmt"

	deck "github.com/gopher-training/deck_task/deck"
)

func main() {
	some := deck.NewArguments{Shuffle: true, AddJokers: 2, Filter: map[deck.Value]struct{}{deck.TWO: {}}, DecksNumber: 2}
	fmt.Println(some)
	customDeck := deck.New(some)
	for _, v := range customDeck {
		fmt.Println(v.ToStringRepresentation())
	}
	customDeck, card := deck.PullRandomCard(customDeck)
	fmt.Println(len(customDeck), card.ToStringRepresentation())
}
