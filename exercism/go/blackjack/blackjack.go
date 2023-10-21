package blackjack

const (
	Stand            = "S"
	Hit              = "H"
	Split            = "P"
	AutomaticallyWin = "W"
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	cards := map[string]int{
		"ace":   11,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"ten":   10,
		"jack":  10,
		"queen": 10,
		"king":  10,
	}
	return cards[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	if card1 == "ace" && card2 == "ace" {
		return Split
	}

	hand := ParseCard(card1) + ParseCard(card2)
	dealer := ParseCard(dealerCard)

	if hand == 21 {
		if dealer < 10 {
			return AutomaticallyWin
		} else {
			return Stand
		}
	}

	if hand >= 17 {
		return Stand
	}

	if hand >= 12 {
		if dealer >= 7 {
			return Hit
		} else {
			return Stand
		}
	}

	return Hit
}
