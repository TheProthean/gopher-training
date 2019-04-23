package main

import (
	"fmt"

	"github.com/gopher-training/deck_task/deck"
)

func main() {
	args := deck.NewArguments{DecksNumber: 1, Shuffle: true}
	playGame(deck.New(args))
}

func playGame(playingDeck []deck.Card) {
	playing := true
	for playing {
		playRound(playingDeck)
		fmt.Println("Hit \"y\", if you want another round.")
		var answer rune
		fmt.Scanf("%c", &answer)
		if answer != 'y' {
			fmt.Println("Closing...")
			playing = false
		}
	}
}

func playRound(playingDeck []deck.Card) {
	fmt.Println("New round.")
	var card deck.Card
	dealerCards := make([]deck.Card, 2)
	playerCards := make([]deck.Card, 2)
	for i := 0; i < 2; i++ {
		playingDeck, card = deck.PullRandomCard(playingDeck)
		dealerCards[i] = card
		playingDeck, card = deck.PullRandomCard(playingDeck)
		playerCards[i] = card
	}
	fmt.Println("Dealer card: ", dealerCards[0].ToStringRepresentation())
	fmt.Println(fmt.Sprintf("Your cards: %s and %s", playerCards[0].ToStringRepresentation(), playerCards[1].ToStringRepresentation()))
	playingDeck, playerCards, playerOverdraw := playerTurn(playingDeck, playerCards)
	playingDeck, dealerCards, dealerOverdraw := dealerTurn(playingDeck, dealerCards)
	playerScore := countScore(playerCards)
	dealerScore := countScore(dealerCards)
	if playerOverdraw {
		fmt.Println("You lost.")
		return
	}
	if dealerOverdraw {
		fmt.Println("You won.")
		return
	}
	if playerScore < dealerScore {
		fmt.Println("You lost.")
		return
	}
	if playerScore > dealerScore {
		fmt.Println("You won.")
		return
	}
	if playerScore < dealerScore {
		fmt.Println("That's a draw.")
		return
	}
}

func countScore(hand []deck.Card) int {
	aceCount := 0
	score := 0
	for _, v := range hand {
		if v.Value == deck.ACE {
			aceCount++
			continue
		}
		if v.Value < deck.JACK {
			score += int(v.Value)
			continue
		}
		if v.Value >= deck.JACK {
			score += 10
			continue
		}
	}
	if score > 10 {
		score += aceCount
	}
	if score <= 10 {
		if aceCount > 0 {
			score += 11 + aceCount - 1
		}
	}
	return score
}

func playerTurn(playingDeck []deck.Card, playerCards []deck.Card) ([]deck.Card, []deck.Card, bool) {
	fmt.Println("Player's turn.")
	hitting := true
	overdraw := false
	var card deck.Card
	for hitting {
		fmt.Printf("Your current score is %d. Do you want to hit or stand?(h/s): ", countScore(playerCards))
		var answer rune
		fmt.Scanf("%c", &answer)
		switch answer {
		case 'h':
			playingDeck, card = deck.PullRandomCard(playingDeck)
			fmt.Printf("You got %s\n", card.ToStringRepresentation())
			playerCards = append(playerCards, card)
			newScore := countScore(playerCards)
			if newScore > 21 {
				fmt.Printf("You overdraw. Your score(%d) is more than 21.\n", newScore)
				overdraw = true
			} else {
				fmt.Printf("Your new score is %d.\n", newScore)
			}
			break
		case 's':
			hitting = false
			break
		default:
			fmt.Println("Wrong choice.")
		}
	}
	return playingDeck, playerCards, overdraw
}

func dealerTurn(playingDeck []deck.Card, dealerCards []deck.Card) ([]deck.Card, []deck.Card, bool) {
	dealerScore := countScore(dealerCards)
	overdraw := false
	var card deck.Card
	fmt.Printf("Dealer's turn. His second card was %s. His score is %d.\n", dealerCards[1].ToStringRepresentation(), dealerScore)
	for dealerNeedDraw(dealerScore, dealerCards) {
		playingDeck, card = deck.PullRandomCard(playingDeck)
		dealerCards = append(dealerCards, card)
		dealerScore = countScore(dealerCards)
		fmt.Printf("Dealer drew %s. His current score is %d.\n", card.ToStringRepresentation(), dealerScore)
	}
	if dealerScore > 21 {
		fmt.Println("Overdraw!")
		overdraw = true
	}
	return playingDeck, dealerCards, overdraw
}

func dealerNeedDraw(dealerScore int, dealerCards []deck.Card) bool {
	return dealerScore <= 16 || (dealerScore == 17 && (dealerCards[0].Value == deck.ACE || dealerCards[1].Value == deck.ACE))
}
