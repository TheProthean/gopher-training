package main

import (
	"fmt"

	"github.com/gopher-training/deck_task/deck"
)

func main() {
	some := deck.NewArguments{Shuffle: true, AddJokers: 2, Filter: map[deck.Value]struct{}{deck.TWO: {}}, DecksNumber: 2}
	fmt.Println(some)
	deck := deck.New(some)
	for _, v := range deck {
		fmt.Println(v)
	}
}
