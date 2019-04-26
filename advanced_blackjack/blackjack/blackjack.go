package blackjack

import (
	"fmt"

	"github.com/gopher-training/deck_task/deck"
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

//RoundState is a struct, that shows AI state of the game, when it's supposed to make a decision
type RoundState struct {
	AIBet             int
	AICards           []deck.Card
	OtherPlayersCards [][]deck.Card
	DealerCard        deck.Card
	allDealerCards    []deck.Card
}

//RoundResult is a struct where stats per Round of our AI are stored
type RoundResult struct {
	IsAWinOrDraw bool
	CashDiff     int
	AIScore      int
	DealerScore  int
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
	TotalRoundsWon    int
	TotalCash         int
	RoundResults      []RoundResult
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
	if bettedCash > g.AIResults.TotalCash {
		bettedCash = g.AIResults.TotalCash
	}
	if bettedCash%10 != 0 {
		bettedCash -= bettedCash % 10
	}
	g.AIResults.TotalCash -= bettedCash
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
		AIBet:             bettedCash,
		AICards:           aiCards,
		OtherPlayersCards: dumbAIsCards,
		DealerCard:        dealerCards[0],
		allDealerCards:    dealerCards,
	}
	for i := 0; i < g.dumbAINumber; i++ {
		g.dumbAITurn(i)
	}
	AINatural := g.aiTurn(smartass)
	dealerNatural := g.dealerTurn()
	g.roundResults(AINatural, dealerNatural)
}

func (g *Game) dealerTurn() bool {
	dealerScore := CountScore(g.currentState.allDealerCards)
	if dealerScore == 21 {
		return true
	}
	var card deck.Card
	for dealerNeedDraw(dealerScore, g.currentState.allDealerCards) {
		g.gameDeck, card = deck.PullFirstCard(g.gameDeck)
		g.currentState.allDealerCards = append(g.currentState.allDealerCards, card)
		dealerScore = CountScore(g.currentState.allDealerCards)
	}
	return false
}

func (g *Game) dumbAITurn(dumbAIID int) {
	AIScore := CountScore(g.currentState.OtherPlayersCards[dumbAIID])
	var card deck.Card
	for AIScore <= 14 {
		g.gameDeck, card = deck.PullFirstCard(g.gameDeck)
		g.currentState.OtherPlayersCards[dumbAIID] = append(g.currentState.OtherPlayersCards[dumbAIID], card)
		AIScore = CountScore(g.currentState.OtherPlayersCards[dumbAIID])
	}
}

func (g *Game) roundResults(AINatural bool, dealerNatural bool) {
	g.AIResults.TotalRoundsPlayed++
	if AINatural && dealerNatural {
		g.AIResults.TotalCash += g.currentState.AIBet
		g.AIResults.RoundResults = append(g.AIResults.RoundResults, RoundResult{true, 0, 21, 21})
	} else if AINatural {
		g.AIResults.TotalCash += int(float64(g.currentState.AIBet) * 2.5)
		g.AIResults.RoundResults = append(g.AIResults.RoundResults, RoundResult{true, int(float64(g.currentState.AIBet) * 1.5), 21, CountScore(g.currentState.allDealerCards)})
		g.AIResults.TotalRoundsWon++
	} else {
		g.AIResults.RoundResults = append(g.AIResults.RoundResults, RoundResult{false, g.currentState.AIBet, CountScore(g.currentState.AICards), 21})
	}
	AIScore := CountScore(g.currentState.AICards)
	dealerScore := CountScore(g.currentState.allDealerCards)
	if AIScore > dealerScore {
		g.AIResults.TotalCash += g.currentState.AIBet * 2
		g.AIResults.RoundResults = append(g.AIResults.RoundResults, RoundResult{true, g.currentState.AIBet, AIScore, dealerScore})
		g.AIResults.TotalRoundsWon++
	}
	if AIScore < dealerScore {
		g.AIResults.RoundResults = append(g.AIResults.RoundResults, RoundResult{false, g.currentState.AIBet, AIScore, dealerScore})
	}
	if AIScore == dealerScore {
		g.AIResults.TotalCash += g.currentState.AIBet
		g.AIResults.RoundResults = append(g.AIResults.RoundResults, RoundResult{true, 0, AIScore, dealerScore})
	}
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
		TotalRoundsWon:    0,
		TotalCash:         settings.InitialCash,
		RoundResults:      []RoundResult{},
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
