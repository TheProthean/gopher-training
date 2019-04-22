package main

import (
	"fmt"

	"github.com/gopher-training/deck_task/deck"
)

func main() {
	some := deck.NewArguments{DecksNumber: 1}
	fmt.Println(some)
	deck := deck.New(some)
	for _, v := range deck {
		fmt.Println(v)
	}
}
