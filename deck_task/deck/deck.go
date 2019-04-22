package deck

//Value is separated from int for better API
type Value int

//Suit is separated from int for better API
type Suit int

const (
	//ACE Value: Ace
	ACE Value = iota
	//TWO Value: 2
	TWO = iota
	//THREE Value: 3
	THREE = iota
	//FOUR Value: 4
	FOUR = iota
	//FIVE Value: 5
	FIVE = iota
	//SIX Value: 6
	SIX = iota
	//SEVEN Value: 7
	SEVEN = iota
	//EIGHT Value: 8
	EIGHT = iota
	//NINE Value: 9
	NINE = iota
	//TEN Value: 10
	TEN = iota
	//JACK Value: Jack
	JACK = iota
	//QUEEN Value: Queen
	QUEEN = iota
	//KING Value: King
	KING = iota
	//JOKER Value: Joker
	JOKER = iota
)

const (
	//HEARTS Suit: Heart
	HEARTS Suit = iota
	//CLUBS Suit: Club
	CLUBS = iota
	//DIAMONDS Suit: Diamond
	DIAMONDS = iota
	//SPADES Suit: Spade
	SPADES = iota
)

//Card is a main type for our deck package that represends a playing card
type Card struct {
	Value Value
	Suit  Suit
}

//NewArguments is a struct to use as optional arguments container in our New function
type NewArguments struct {
	SortingFunc func(i, j Card) bool
	//WARNING: shuffle overrides all other sorting arguments
	Shuffle     bool
	AddJokers   int
	Filter      []Value
	DecksNumber int
}

//New is our function for creating new deck of playing cards
func New(args NewArguments) []Card {
	deck := make([]Card, 52, 52)
	for i := HEARTS; i <= SPADES; i++ {
		for j := ACE; j <= KING; j++ {
			deck[int(i)*13+int(j)] = Card{Value: j, Suit: i}
		}
	}
	return deck
}
