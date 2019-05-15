package blackjack

import (
	"fmt"
	"os"

	"github.com/gopher-training/deck_task/deck"
)

//AIDecision is a secured type for our AIs decisions generalization
type AIDecision int

const (
	//HIT will be counted as if AI wants a hit
	HIT AIDecision = iota
	//STAND will be counted as if AI wants to stop hitting
	STAND
	//SPLIT will be counted as if AI wants to split his cards
	SPLIT
	//DOUBLEDOWN will be counted as if AI wants to take a doubledown option
	DOUBLEDOWN
)

//AI is our AI interface, that will be playing our game
type AI interface {
	//This function should return one of 4 AIDecision values
	//It should return SPLIT or DOUBLEDOWN ONLY when flag is TRUE, otherwise the game will end immediatly
	MakeDecision(state RoundState, AICash int, expandedOptions bool) AIDecision
	//This function should return amount of cash, that AI wants to bet
	//Bet amount should be divisible by 10, otherwise it will be floored to closest number divisible by 10
	//If it is more, than what AI has, AI will bet all his cash
	//AI should bet at least 10 cash, if he has less than 10 cash the game stops
	BetCash(AICash int) int
}

func (g *Game) aiTurn(smartass AI, firstTurnExpanded bool) (bool, *Game) {
	var card deck.Card
	var splitHand []deck.Card
	splitOccured := false
	AIturn := true
	firstIteration := firstTurnExpanded
	AIScore := CountScore(g.currentState.AICards)
	if AIScore == 21 {
		return true, nil
	}
	if len(g.currentState.AICards) == 1 && g.currentState.AICards[0].Value == deck.ACE {
		g.gameDeck, card = deck.PullFirstCard(g.gameDeck)
		g.currentState.AICards = append(g.currentState.AICards, card)
		return false, nil
	}
	for AIturn {
		AIDecision := smartass.MakeDecision(g.currentState, g.AIResults.TotalCash, firstIteration)
		if (AIDecision == SPLIT || AIDecision == DOUBLEDOWN) && !firstIteration {
			fmt.Printf("Game has stopped - AI chose forbidden option.\nHe chose to split or to doubledown when it was not allowed.\n")
			os.Exit(1)
		}
		firstIteration = false
		switch AIDecision {
		case HIT:
			g.gameDeck, card = deck.PullFirstCard(g.gameDeck)
			g.currentState.AICards = append(g.currentState.AICards, card)
			break
		case STAND:
			AIturn = false
			break
		case SPLIT:
			if g.currentState.AICards[0].Value != g.currentState.AICards[1].Value {
				fmt.Printf("Game has stopped - AI chose forbidden option.\nWhile not having 2 cards with the same value he chose to split.\n")
				os.Exit(1)
			}
			if g.AIResults.TotalCash < g.currentState.AIBet {
				fmt.Printf("Game has stopped - AI chose forbidden option.\nWith total cash of %d and bet of %d he chose to split.\n", g.AIResults.TotalCash, g.currentState.AIBet)
				os.Exit(1)
			}
			g.AIResults.TotalCash -= g.currentState.AIBet
			splitOccured = true
			splitHand = []deck.Card{g.currentState.AICards[1]}
			g.currentState.AICards = []deck.Card{g.currentState.AICards[0]}
			break
		case DOUBLEDOWN:
			AIScore = CountScore(g.currentState.AICards)
			if AIScore <= 8 || AIScore >= 12 {
				fmt.Printf("Game has stopped - AI chose forbidden option.\nWith a score of %d he chose to doubledown.\n", AIScore)
				os.Exit(1)
			}
			if g.AIResults.TotalCash < g.currentState.AIBet {
				fmt.Printf("Game has stopped - AI chose forbidden option.\nWith total cash of %d and bet of %d he chose to doubledown.\n", g.AIResults.TotalCash, g.currentState.AIBet)
				os.Exit(1)
			}
			g.AIResults.TotalCash -= g.currentState.AIBet
			g.currentState.AIBet *= 2
			g.gameDeck, card = deck.PullFirstCard(g.gameDeck)
			g.currentState.AICards = append(g.currentState.AICards, card)
			AIturn = false
			break
		}
	}
	if splitOccured {
		splitResult := Results{
			TotalRoundsPlayed: 0,
			TotalRoundsWon:    0,
			TotalCash:         0,
			RoundResults:      []RoundResult{},
		}
		splitRoundState := RoundState{
			AIBet:             g.currentState.AIBet,
			AICards:           splitHand,
			OtherPlayersCards: append(g.currentState.OtherPlayersCards, g.currentState.AICards),
			DealerCard:        g.currentState.DealerCard,
			allDealerCards:    g.currentState.allDealerCards,
		}
		splitGame := Game{
			gameDeck:     g.gameDeck,
			currentState: splitRoundState,
			AIResults:    splitResult,
		}
		return false, &splitGame
	}
	return false, nil
}
