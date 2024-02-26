package blackjack

var cards = map[string]int{
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

func ParseCard(card string) int {
	return cards[card]
}

func FirstTurn(card1, card2, dealerCard string) string {
	sum := cards[card1] + cards[card2]

	switch {
	case sum == 22:
		return "p"
	case sum == 21:
		if cards[dealerCard] < 10 {
			return "W"
		}

		return "S"
	case sum >= 17 && sum <= 20:
		return "S"

	case sum >= 12 && sum <= 16:
		if cards[dealerCard] >= 7 {
			return "H"
		}

		return "S"

	default:
		return "H"
	}
}
