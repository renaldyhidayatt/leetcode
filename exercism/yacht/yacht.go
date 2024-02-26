package yacht

import (
	"log"
	"sort"
)

var (
	littleStraight = []int{1, 2, 3, 4, 5}
	bigStraight    = []int{2, 3, 4, 5, 6}
)

func Score(dice []int, category string) int {
	sort.Ints(dice)

	var sum int

	counts := map[int]int{}
	sums := map[int]int{}

	for _, die := range dice {
		sum += die
		sums[die] += die
		counts[die]++
	}

	equal := func(arr []int) bool {
		for i, v := range arr {
			if dice[i] != v {
				return false
			}
		}

		return true
	}

	switch category {
	case "big straight":
		if equal(bigStraight) {
			return 30
		}

	case "choice":
		return sum

	case "fives":
		return sums[5]
	case "four of a kind":
		if counts[dice[0]] >= 4 {
			return dice[0] * 4
		}

		if counts[dice[1]] >= 4 {
			return dice[1] * 4
		}
	case "fours":
		return sums[4]
	case "full house":
		if (counts[dice[0]] == 3 && counts[dice[4]] == 2) || (counts[dice[0]] == 2 && counts[dice[4]] == 3) {
			return sum
		}
	case "little straight":
		if equal(littleStraight) {
			return 30
		}
	case "ones":
		return sums[1]
	case "sixes":
		return sums[6]
	case "threes":
		return sums[3]

	case "twos":
		return sums[2]
	case "yacht":
		if sum == dice[0]*5 {
			return 50
		}
	default:
		log.Fatalf("unknown category: %v", category)
	}

	return 0

}
