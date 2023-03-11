package numberofvalidwordsforeach

func findNumOfValidWords(words []string, puzzles []string) []int {
	wordBitStatusMap, res := make(map[uint32]int, 0), []int{}
	for _, w := range words {
		wordBitStatusMap[toBitMap([]byte(w))]++
	}

	for _, p := range puzzles {
		var bitMap uint32
		var totalNum int
		bitMap |= (1 << (p[0] - 'a'))
		findNum([]byte(p)[1:], bitMap, &totalNum, wordBitStatusMap)
		res = append(res, totalNum)
	}
	return res
}

func toBitMap(word []byte) uint32 {
	var res uint32

	for _, b := range word {
		res |= (1 << (b - 'a'))
	}
	return res
}

func findNum(puzzles []byte, bitMap uint32, totalNum *int, m map[uint32]int) {
	if len(puzzles) == 0 {
		*totalNum = *totalNum + m[bitMap]
		return
	}

	findNum(puzzles[1:], bitMap, totalNum, m)

	bitMap |= (1 << (puzzles[0] - 'a'))

	findNum(puzzles[1:], bitMap, totalNum, m)

	bitMap ^= (1 << (puzzles[0] - 'a'))
	return
}
