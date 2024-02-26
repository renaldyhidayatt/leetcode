package moserdebruijnsequence

func MoserDeBruijnSequence(num int) []int {
	sequence := []int{}

	for i := 0; i < num; i++ {
		res := generateNthTerm(i)

		sequence = append(sequence, res)
	}

	return sequence
}

func generateNthTerm(num int) int {
	if num == 0 || num == 1 {
		return num
	}

	if num%2 == 0 {
		return 4 * generateNthTerm(num/2)
	}

	return 4*generateNthTerm(num/2) + 1
}
