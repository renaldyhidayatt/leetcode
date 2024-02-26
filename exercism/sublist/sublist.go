package sublist

type Relation string

const (
	RelationEqual     Relation = "equal"
	RelationSublist   Relation = "sublist"
	RelationSuperlist Relation = "superlist"
	RelationUnequal   Relation = "unequal"
)

func Sublist(s1, s2 []int) Relation {
	if contains(s1, s2) {
		if len(s1) == len(s2) {
			return RelationEqual
		}

		return RelationSuperlist
	} else if contains(s2, s1) {
		return RelationSublist
	}

	return RelationUnequal
}

func contains(s1, s2 []int) bool {
	i1, i2 := 0, 0

	for ; i1 < len(s1) && i2 < len(s2); i1++ {
		if s1[i1] == s2[i2] {
			i2++
		} else {
			i1, i2 = i1-i2, 0
		}
	}

	return i2 == len(s2)
}
