package on399

var (
	res      []string
	mp       map[string]int
	n        int
	length   int
	maxScore int
	str      string
)

func Backtrace(i int, cur string, lmoves int, rmoves int, score int) {
	if lmoves < 0 || rmoves < 0 || score < 0 || score > maxScore {
		return
	}

	if lmoves == 0 && rmoves == 0 {
		if len(cur) == length {
			if _, ok := mp[cur]; !ok {
				res = append(res, cur)

				mp[cur] = 1
			}

			return
		}
	}

	if i == n {
		return
	}
	if str[i] == '(' {
		Backtrace(i+1, cur+string('('), lmoves, rmoves, score+1)
		Backtrace(i+1, cur, lmoves-1, rmoves, score)
	} else if str[i] == ')' {
		Backtrace(i+1, cur+string(')'), lmoves, rmoves, score-1)
		Backtrace(i+1, cur, lmoves, rmoves-1, score)
	} else {
		Backtrace(i+1, cur+string(str[i]), lmoves, rmoves, score)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
