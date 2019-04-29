package main

import (
	"fmt"

	"github.com/gopher-training/advanced_blackjack/blackjack"
	"github.com/gopher-training/deck_task/deck"
)

//AI ...
type AI struct {
}

//MakeDecision ...
func (ai *AI) MakeDecision(state blackjack.RoundState, AICash int, expandedOptions bool) blackjack.AIDecision {
	AIScore := blackjack.CountScore(state.AICards)
	if expandedOptions {
		if shouldDoubleDown(AIScore, state.DealerCard, AICash, state.AIBet) {
			return blackjack.DOUBLEDOWN
		}
		if shouldSplit(state.AICards, AICash, state.AIBet) {
			return blackjack.SPLIT
		}
		if shouldHitWithSoftHand(state.AICards, AIScore) {
			return blackjack.HIT
		}
	}
	if shouldHit(AIScore, state.DealerCard) {
		return blackjack.HIT
	}
	return blackjack.STAND
}

func shouldDoubleDown(score int, dealerCard deck.Card, AICash int, AIBet int) bool {
	return score == 11 || (score == 10 && !(dealerCard.Value >= deck.JACK || dealerCard.Value == deck.ACE)) || (score == 9 && !(dealerCard.Value >= deck.SEVEN || dealerCard.Value == deck.ACE)) && AIBet <= AICash
}

func shouldSplit(hand []deck.Card, AICash int, AIBet int) bool {
	return hand[0].Value == hand[1].Value && (hand[0].Value == deck.ACE || hand[0].Value == deck.EIGHT) && AIBet <= AICash
}

func shouldHitWithSoftHand(hand []deck.Card, score int) bool {
	return (hand[0].Value == deck.ACE || hand[1].Value == deck.ACE) && score < 18
}

func shouldHit(score int, dealerCard deck.Card) bool {
	return (score < 17 && (dealerCard.Value >= deck.SEVEN || dealerCard.Value == deck.ACE)) || (score < 13 && (dealerCard.Value <= deck.SIX && dealerCard.Value > deck.ACE))
}

//BetCash ...
func (ai *AI) BetCash(AICash int) int {
	if AICash >= 50 {
		return 20
	}
	return 10
}

func main() {
	gameSettings := blackjack.GameOptions{
		GameCount:    20,
		DeckCount:    3,
		DumbAINumber: 2,
		InitialCash:  300,
	}
	newAI := AI{}
	game := blackjack.New(gameSettings)
	game.Play(&newAI)
	fmt.Printf("AI played %d games.\nHe won in %d of them.\nHis total cash at the end of the game is %d.\n", game.AIResults.TotalRoundsPlayed, game.AIResults.TotalRoundsWon, game.AIResults.TotalCash)
	fmt.Println(game.AIResults.RoundResults)
}
