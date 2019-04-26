package deck

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//Value is separated from int for better API
type Value int

//Suit is separated from int for better API
type Suit int

const (
	//ACE Value: Ace
	ACE Value = iota
	//TWO Value: 2
	TWO
	//THREE Value: 3
	THREE
	//FOUR Value: 4
	FOUR
	//FIVE Value: 5
	FIVE
	//SIX Value: 6
	SIX
	//SEVEN Value: 7
	SEVEN
	//EIGHT Value: 8
	EIGHT
	//NINE Value: 9
	NINE
	//TEN Value: 10
	TEN
	//JACK Value: Jack
	JACK
	//QUEEN Value: Queen
	QUEEN
	//KING Value: King
	KING
	//JOKER Value: Joker
	JOKER
)

const (
	//HEARTS Suit: Heart
	HEARTS Suit = iota
	//CLUBS Suit: Club
	CLUBS
	//DIAMONDS Suit: Diamond
	DIAMONDS
	//SPADES Suit: Spade
	SPADES
)

var valueNamesMap = map[Value]string{
	ACE: "Ace", TWO: "2", THREE: "3",
	FOUR: "4", FIVE: "5", SIX: "6",
	SEVEN: "7", EIGHT: "8", NINE: "9",
	TEN: "10", JACK: "Jack", QUEEN: "Queen",
	KING: "King", JOKER: "Joker",
}

var suitNamesMap = map[Suit]string{
	HEARTS: "Hearts", CLUBS: "Clubs",
	DIAMONDS: "Diamonds", SPADES: "Spades",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

//Card is a main type for our deck package that represends a playing card
type Card struct {
	Value Value
	Suit  Suit
}

//ToStringRepresentation is a function to convert our Card type to human-readable condition
func (c Card) ToStringRepresentation() string {
	return fmt.Sprintf("%s of %s", valueNamesMap[c.Value], suitNamesMap[c.Suit])
}

//NewArguments is a struct to use as optional arguments container in our New function
//Be aware, that Shuffle argument overrides SortingFunc argument
type NewArguments struct {
	//Function that defines order of cards in new deck
	SortingFunc func(i, j int) bool
	/*Flag that defines, will the order of cards be random or not
	WARNING: shuffle overrides all other sorting arguments*/
	Shuffle bool
	/*Amount of Jokers to add to new deck.
	Jokers are added with Suit property in exactly same order suits are in default deck*/
	AddJokers int
	//This argument defines cards, values of which we don't want to see in new deck
	Filter map[Value]struct{}
	//This argument defines how much usual decks should new deck contain
	DecksNumber int
}

//New is our function for creating new deck of playing cards.
//By default it is sorted in order: Suits - Hearts, Clubs, Diamonds, Spades;
//Values - Ace, 2, 3, 4, 5, 6, 7, 8, 9, 10, Jack, Queen, King.
func New(args NewArguments) []Card {
	deck := []Card{}
	if args.DecksNumber == 0 {
		args.DecksNumber = 1
	}
	for k := 0; k < args.DecksNumber; k++ {
		for i := HEARTS; i <= SPADES; i++ {
			for j := ACE; j <= KING; j++ {
				newCard := Card{Value: j, Suit: i}
				if _, in := args.Filter[newCard.Value]; !in {
					deck = append(deck, newCard)
				}
			}
		}
	}
	if args.AddJokers != 0 {
		for k := 0; k < args.AddJokers; k++ {
			deck = append(deck, Card{Value: JOKER, Suit: Suit(k % 4)})
		}
	}
	if args.Shuffle {
		rand.Shuffle(len(deck), func(i, j int) {
			deck[i], deck[j] = deck[j], deck[i]
		})
		return deck
	}
	if args.SortingFunc != nil {
		sort.SliceStable(deck, args.SortingFunc)
	}
	return deck
}

//PullRandomCard pulls one random card from deck. Returns updated deck and pulled card.
func PullRandomCard(deck []Card) ([]Card, Card) {
	cardNum := rand.Intn(len(deck) - 1)
	pulledCard := deck[cardNum]
	updatedDeck := make([]Card, len(deck)-1)
	copy(updatedDeck, deck[:cardNum])
	copy(updatedDeck[cardNum:], deck[cardNum+1:])
	return updatedDeck, pulledCard
}

//PutCardBackInDeck puts card back into deck, checking that this card is not in the deck already(so it is only for single deck).
//Jokers though won't be put in the deck anyway, even if there is no jokers in the deck.
func PutCardBackInDeck(deck []Card, card Card) ([]Card, error) {
	if card.Value == JOKER {
		return deck, errors.New("Can't put jokers in the deck. Create new deck if you want to use jokers.")
	}
	for _, v := range deck {
		if v == card {
			return deck, errors.New("This card is already in the deck")
		}
	}
	deck = append(deck, card)
	return deck, nil
}

//PutCardBackInDeckUnsafe is an unsafe variant of function PutCardBackInDeck. Use it when deck consists of multiple decks.
func PutCardBackInDeckUnsafe(deck []Card, card Card) []Card {
	deck = append(deck, card)
	return deck
}

//Shuffle is a function that allows us to separately shuffle already existing deck. Returns shuffled deck.
func Shuffle(deck []Card) []Card {
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	return deck
}

//PullFirstCard pulls one card from the top of the deck. Returns updated deck and pulled card.
func PullFirstCard(deck []Card) ([]Card, Card) {
	pulledCard := deck[0]
	updatedDeck := make([]Card, len(deck)-1)
	copy(updatedDeck, deck[1:])
	return updatedDeck, pulledCard
}
