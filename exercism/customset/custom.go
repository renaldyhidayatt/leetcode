package customset

import "fmt"

type Set map[string]struct{}

func New() Set {
	return make(Set)
}

func NewFromSlice(l []string) Set {
	result := New()

	for _, s := range l {
		result.Add(s)
	}

	return result
}

func (s Set) String() string {
	var i int

	x := len(s)

	result := "{"

	for k := range s {
		i++

		result += fmt.Sprintf("\"%v\"", k)
		if i < x {
			result += ", "
		}
	}

	result += "}"

	return result
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	_, ok := s[elem]

	return ok
}

func (s Set) Add(elem string) {
	s[elem] = struct{}{}
}

func Subset(s1, s2 Set) bool {
	for k := range s1 {
		if !s2.Has(k) {
			return false
		}
	}

	return true
}

func Disjoint(s1, s2 Set) bool {
	return len(Intersection(s1, s2)) == 0
}

func Equal(s1, s2 Set) bool {
	return Subset(s1, s2) && len(s1) == len(s2)
}

func Intersection(s1, s2 Set) Set {
	result := New()

	for k := range s1 {
		if s2.Has(k) {
			result.Add(k)
		}
	}
	return result
}

func Difference(s1, s2 Set) Set {
	result := New()

	for s := range s1 {
		if !s2.Has(s) {
			result.Add(s)
		}
	}

	return result
}

func Union(s1, s2 Set) Set {
	for s := range s1 {
		s2.Add(s)
	}

	return s2
}
