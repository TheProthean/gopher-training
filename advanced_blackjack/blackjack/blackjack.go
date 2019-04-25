package blackjack

import (
	"fmt"

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

//GameOptions is a struct that stores game settings.
//WARNING: AI should have at least 10 cash at the start of the game
//WARNING: card amount should exceed number of AIs * 2 + 2 at least by 40
type GameOptions struct {
	GameCount    int
	DeckCount    int
	DumbAINumber int
	InitialCash  int
}

//AI is our AI interface, that will be playing our game
type AI interface {
	//This function should return one of 4 AIDecision values
	//It should return SPLIT or DOUBLEDOWN ONLY when flag is TRUE, otherwise the game will end immediatly
	MakeDecision(state RoundState, expandedOptions bool) AIDecision
	//This function should return amount of cash, that AI wants to bet
	//If it is more, than what AI has, AI will bet all his cash
	//AI should bet at least 10 cash, if he has less than 10 cash the game stops
	BetCash(AICash int) int
}

//RoundState is a struct, that shows AI state of the game, when it's supposed to make a decision
type RoundState struct {
	AICards        []deck.Card
	DumbAiCards    [][]deck.Card
	DealerCard     deck.Card
	allDealerCards []deck.Card
}

//Game is our game, that our AI or Player will be playing.
type Game struct {
	gameDeck     []deck.Card
	dumbAINumber int
	roundsLeft   int
	currentState RoundState
	AIResults    Results
}

//Results is a storage of our AI stats
type Results struct {
	TotalRoundsPlayed int
	TotalCash         int
	Losses            []struct{ int int }
	Wins              []struct{ int int }
}

//Play is a main function of our Game struct, that starts the game
func (g *Game) Play(smartass AI) {
	if g.dumbAINumber*2+2+40 >= len(g.gameDeck) {
		fmt.Println("This game cannot be played - there is not enough cards.")
		return
	}
	if g.AIResults.TotalCash < 10 {
		fmt.Println("This game cannot be played - AI should have at least 10 cash ")
	}
	for g.roundsLeft > 0 && g.AIResults.TotalCash >= 10 {
		g.playRound(smartass)
	}
	fmt.Println("The game has concluded.")
}

func (g *Game) playRound(smartass AI) {
	g.gameDeck = deck.Shuffle(g.gameDeck)
	bettedCash := smartass.BetCash(g.AIResults.TotalCash)
	var card deck.Card
	dealerCards := make([]deck.Card, 2)
	aiCards := make([]deck.Card, 2)
	dumbAIsCards := [][]deck.Card{}
	for i := 0; i < 2; i++ {
		for j := 0; j < g.dumbAINumber; j++ {
			if i == 0 {
				dumbAICards := make([]deck.Card, 2)
				dumbAIsCards = append(dumbAIsCards, dumbAICards)
			}
			g.gameDeck, card = deck.PullFirstCard(g.gameDeck)
			dumbAIsCards[j][i] = card
		}
		g.gameDeck, card = deck.PullFirstCard(g.gameDeck)
		aiCards[i] = card
		g.gameDeck, card = deck.PullFirstCard(g.gameDeck)
		dealerCards[i] = card
	}
	g.currentState = RoundState{
		AICards:        aiCards,
		DumbAiCards:    dumbAIsCards,
		DealerCard:     dealerCards[0],
		allDealerCards: dealerCards,
	}
}

func (g *Game) dealerTurn() bool {
	dealerScore := CountScore(g.currentState.allDealerCards)
	overdraw := false
	var card deck.Card
	for dealerNeedDraw(dealerScore, g.currentState.allDealerCards) {
		g.gameDeck, card = deck.PullFirstCard(g.gameDeck)
		g.currentState.allDealerCards = append(g.currentState.allDealerCards, card)
		dealerScore = CountScore(g.currentState.allDealerCards)
	}
	if dealerScore > 21 {
		overdraw = true
	}
	return overdraw
}

func (g *Game) dumbAITurn(dumbAIID int) {
	AIScore := CountScore(g.currentState.DumbAiCards[dumbAIID])
	var card deck.Card
	for AIScore <= 12 {
		g.gameDeck, card = deck.PullFirstCard(g.gameDeck)
		g.currentState.DumbAiCards[dumbAIID] = append(g.currentState.DumbAiCards[dumbAIID], card)
		AIScore = CountScore(g.currentState.DumbAiCards[dumbAIID])
	}
}

func (g *Game) aiTurn(smartass AI) bool {
	AIDecision := smartass.MakeDecision(g.currentState, true)
	return false
}

func dealerNeedDraw(dealerScore int, dealerCards []deck.Card) bool {
	return dealerScore <= 16 || (dealerScore == 17 && (dealerCards[0].Value == deck.ACE || dealerCards[1].Value == deck.ACE))
}

//New is a function that creates new game of blackjack
func New(settings GameOptions) Game {
	deckArguments := deck.NewArguments{
		DecksNumber: settings.DeckCount,
		Shuffle:     true,
	}
	aires := Results{
		TotalRoundsPlayed: 0,
		TotalCash:         settings.InitialCash,
	}
	return Game{
		gameDeck:     deck.New(deckArguments),
		dumbAINumber: settings.DumbAINumber,
		roundsLeft:   settings.GameCount,
		AIResults:    aires,
	}
}

//CountScore is a function, that can be used in AI functions, so that AI can calculate score of his hand
func CountScore(hand []deck.Card) int {
	aceCount := 0
	score := 0
	for _, v := range hand {
		if v.Value == deck.ACE {
			aceCount++
			continue
		}
		if v.Value < deck.JACK {
			score += int(v.Value) + 1
			continue
		}
		score += 10
		continue
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
