package strings

func IsSubsequence(s string, t string) bool{
	if len(s) > len(t){
		return false
	}

	if s == t{
		return true
	}

	if len(s) == 0{
		return true
	}

	sIndex := 0;

	for tIndex := 0; tIndex < len(t); tIndex++{
		if s[sIndex] == t[tIndex]{
			sIndex++
		}

		if sIndex == len(s){
			return true
		}
	}

	return false
}