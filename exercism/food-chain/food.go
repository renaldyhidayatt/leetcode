package foodchain

import "strings"

var words = []struct {
	animal string
	once   string
	repeat string
}{

	{
		"fly",
		"",
		"I don't know why she swallowed the fly. Perhaps she'll die.",
	},
	{
		"spider",
		"It wriggled and jiggled and tickled inside her.",
		"She swallowed the spider to catch the fly.",
	},
	{
		"bird",
		"How absurd to swallow a bird!",
		"She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.",
	},
	{
		"cat",
		"Imagine that, to swallow a cat!",
		"She swallowed the cat to catch the bird.",
	},
	{
		"dog",
		"What a hog, to swallow a dog!",
		"She swallowed the dog to catch the cat.",
	},
	{
		"goat",
		"Just opened her throat and swallowed a goat!",
		"She swallowed the goat to catch the dog.",
	},
	{
		"cow",
		"I don't know how she swallowed a cow!",
		"She swallowed the cow to catch the goat.",
	},
	{
		"horse",
		"She's dead, of course!",
		"",
	},
}

func Verse(v int) string {
	verse := []string{"I know an old lady who swallowed a " + words[v-1].animal + "."}

	if v > 1 {
		verse = append(verse, words[v-1].once)
	}

	for i := v - 1; v < 8 && i >= 0; i-- {
		verse = append(verse, words[i].repeat)
	}

	return strings.Join(verse, "\n")
}

func Verses(start, end int) string {
	var verses []string

	for i := start; i <= end; i++ {
		verses = append(verses, Verse(i))
	}

	return strings.Join(verses, "\n\n")
}

func Song() string {
	return Verses(1, 8)
}
