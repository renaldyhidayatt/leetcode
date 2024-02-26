package listops

type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	acc := initial

	for _, val := range s {
		acc = fn(acc, val)
	}

	return acc
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	acc := initial

	for i := len(s) - 1; i >= 0; i-- {
		acc = fn(s[i], acc)
	}

	return acc
}

func (s IntList) Filter(fn func(int) bool) IntList {
	res := IntList{}

	for _, val := range s {
		if fn(val) {
			res = append(res, val)
		}
	}

	return res
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	res := IntList{}

	for _, val := range s {
		res = append(res, fn(val))
	}

	return res
}

func (s IntList) Reverse() IntList {
	res := IntList{}

	for i := len(s) - 1; i >= 0; i-- {
		res = append(res, s[i])
	}

	return res
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	res := s

	for _, list := range lists {
		res = append(res, list...)
	}

	return res
}
