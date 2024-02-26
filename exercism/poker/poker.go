package poker

import (
	"errors"
	"sort"
	"strings"
)

type Card struct {
	value int
	suit  string
}

type Hand struct {
	cards     [5]Card
	rank      int
	original  string
	multiples [5]int
}

var faceToValue = map[string]int{"A": 14, "K": 13, "Q": 12, "J": 11, "10": 10,
	"9": 9, "8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2}

func BestHand(sHands []string) ([]string, error) {
	hands := make([]Hand, 0)

	for _, sHand := range sHands {
		if nh, err := NewHand(sHand); err != nil {
			return sHands, err
		} else {
			hands = append(hands, nh)
		}
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].rank > hands[j].rank
	})

	bestRank := hands[0].rank
	best := make([]string, 0, 1)

	for _, hand := range hands {
		if bestRank == hand.rank {
			best = append(best, hand.original)
		}
	}

	return best, nil
}

func NewHand(hand string) (Hand, error) {
	sCards := strings.Fields(hand)

	if len(sCards) != 5 {
		return Hand{}, errors.New("requires five cards")
	}

	newHand := Hand{original: hand}

	for i, sCard := range sCards {
		card, err := newCard([]rune(sCard))

		if err != nil {
			return newHand, err
		}

		newHand.cards[i] = card
	}

	faceValue := make([]int, len(faceToValue)+5)

	for _, card := range newHand.cards {
		faceValue[card.value]++
	}

	for _, v := range faceValue {
		if v > 0 {
			newHand.multiples[v]++
		}
	}

	sort.Slice(newHand.cards[:], func(i, j int) bool {
		return newHand.cards[i].value < newHand.cards[j].value
	})

	allCards := newHand.cards[0].value + newHand.cards[1].value<<4 + newHand.cards[2].value<<8 +
		newHand.cards[3].value<<12 + newHand.cards[4].value<<16

	switch {
	case newHand.flush() && newHand.straight(): // Royal Flush
		newHand.rank = 8<<24 + newHand.cards[3].value<<20
	case newHand.multiples[4] == 1: // Four of a Kind
		newHand.rank = 7<<24 + newHand.cards[2].value<<20 + allCards
	case newHand.multiples[2] == 1 && newHand.multiples[3] == 1: // Full House
		newHand.rank = 6<<24 + newHand.cards[2].value<<20 + allCards
	case newHand.flush(): // Flush
		newHand.rank = 5<<24 + allCards
	case newHand.straight() && newHand.cards[0].value == 2 && newHand.cards[4].value == 14: // Straight - Ace Low
		newHand.rank = 4 << 24
	case newHand.straight(): // Straight
		newHand.rank = 4<<24 + newHand.cards[4].value<<20
	case newHand.multiples[3] == 1: // Three of a Kind
		newHand.rank = 3<<24 + newHand.cards[2].value<<20 + allCards
	case newHand.multiples[2] == 2: // Two Pair
		newHand.rank = 2<<24 + allCards
	case newHand.multiples[2] == 1: // One Pair
		newHand.rank = 1<<24 + allCards
	default: // High Cards
		newHand.rank = allCards
	}

	return newHand, nil
}

func newCard(card []rune) (Card, error) {
	length := len(card)

	if length < 2 || length > 3 {
		return Card{}, errors.New("invalid card")
	}

	suit := string(card[length-1:])
	value, ok := faceToValue[string(card[:length-1])]

	if !ok || !strings.ContainsAny("♤♡♢♧", suit) {
		return Card{}, errors.New("invalid card (suit)")
	}

	return Card{value: value, suit: suit}, nil
}

func (h *Hand) straight() bool {
	return h.cards[0].value+1 == h.cards[1].value &&
		h.cards[1].value+1 == h.cards[2].value &&
		h.cards[2].value+1 == h.cards[3].value &&
		(h.cards[3].value+1 == h.cards[4].value ||
			(h.cards[0].value == 2 && h.cards[4].value == 14))
}

func (h *Hand) flush() bool {
	return h.cards[0].suit == h.cards[1].suit &&
		h.cards[1].suit == h.cards[2].suit &&
		h.cards[2].suit == h.cards[3].suit &&
		h.cards[3].suit == h.cards[4].suit
}
