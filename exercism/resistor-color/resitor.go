package resistorcolor

func Colors() []string {
	return []string{
		"black",
		"brown",
		"red",
		"orange",
		"yellow",
		"green",
		"blue",
		"violet",
		"grey",
		"white",
	}
}

func ColorCode(color string) int {
	for i, element := range Colors() {
		if element == color {
			return i
		}
	}

	return -1
}
