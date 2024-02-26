package resitorcolortrio

import (
	"math"
	"strconv"
)

func Label(colors []string) string {
	colorMap := map[string]int{
		"black":  0,
		"brown":  1,
		"red":    2,
		"orange": 3,
		"yellow": 4,
		"green":  5,
		"blue":   6,
		"violet": 7,
		"grey":   8,
		"white":  9,
	}

	resitorVal := (10*colorMap[colors[0]] + colorMap[colors[1]]) * int(math.Pow10(colorMap[colors[2]]))

	if resitorVal >= 1000000000 {
		return strconv.Itoa(resitorVal/1000000000) + " gigaohms"
	} else if resitorVal >= 1000000 {
		return strconv.Itoa(resitorVal/1000000) + " megaohms"
	} else if resitorVal >= 1000 {
		return strconv.Itoa(resitorVal/1000) + " kiloohms"
	} else {
		return strconv.Itoa(resitorVal) + " ohms"
	}
}
