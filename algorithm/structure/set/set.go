package set

func New[T comparable](items ...T) Set[T] {
	st := set[T]{
		elements: make(map[T]bool),
	}
	for _, item := range items {
		st.Add(item)
	}
	return &st
}

type Set[T comparable] interface {
	Add(item T)
	Delete(item T)
	Len() int
	GetItems() []T
	In(item T) bool
	IsSubsetOf(set2 Set[T]) bool
	IsProperSubsetOf(set2 Set[T]) bool
	IsSupersetOf(set2 Set[T]) bool
	IsProperSupersetOf(set2 Set[T]) bool
	Union(set2 Set[T]) Set[T]
	Intersection(set2 Set[T]) Set[T]
	Difference(set2 Set[T]) Set[T]
	SymmetricDifference(set2 Set[T]) Set[T]
}

type set[T comparable] struct {
	elements map[T]bool
}

func (st *set[T]) Add(value T) {
	st.elements[value] = true
}

func (st *set[T]) Delete(value T) {
	delete(st.elements, value)
}

func (st *set[T]) GetItems() []T {
	keys := make([]T, 0, len(st.elements))
	for k := range st.elements {
		keys = append(keys, k)
	}
	return keys
}

func (st *set[T]) Len() int {
	return len(st.elements)
}

func (st *set[T]) In(value T) bool {
	if _, in := st.elements[value]; in {
		return true
	}
	return false
}

func (st *set[T]) IsSubsetOf(superSet Set[T]) bool {
	if st.Len() > superSet.Len() {
		return false
	}

	for _, item := range st.GetItems() {
		if !superSet.In(item) {
			return false
		}
	}
	return true
}

func (st *set[T]) IsProperSubsetOf(superSet Set[T]) bool {
	if st.Len() == superSet.Len() {
		return false
	}
	return st.IsSubsetOf(superSet)
}

func (st *set[T]) IsSupersetOf(subSet Set[T]) bool {
	return subSet.IsSubsetOf(st)
}

func (st *set[T]) IsProperSupersetOf(subSet Set[T]) bool {
	if st.Len() == subSet.Len() {
		return false
	}
	return st.IsSupersetOf(subSet)
}

func (st *set[T]) Union(st2 Set[T]) Set[T] {
	unionSet := New[T]()
	for _, item := range st.GetItems() {
		unionSet.Add(item)
	}
	for _, item := range st2.GetItems() {
		unionSet.Add(item)
	}
	return unionSet
}

func (st *set[T]) Intersection(st2 Set[T]) Set[T] {
	intersectionSet := New[T]()
	var minSet, maxSet Set[T]
	if st.Len() > st2.Len() {
		minSet = st2
		maxSet = st
	} else {
		minSet = st
		maxSet = st2
	}
	for _, item := range minSet.GetItems() {
		if maxSet.In(item) {
			intersectionSet.Add(item)
		}
	}
	return intersectionSet
}

func (st *set[T]) Difference(st2 Set[T]) Set[T] {
	differenceSet := New[T]()
	for _, item := range st.GetItems() {
		if !st2.In(item) {
			differenceSet.Add(item)
		}
	}
	return differenceSet
}

func (st *set[T]) SymmetricDifference(st2 Set[T]) Set[T] {
	symmetricDifferenceSet := New[T]()
	dropSet := New[T]()
	for _, item := range st.GetItems() {
		if st2.In(item) {
			dropSet.Add(item)
		} else {
			symmetricDifferenceSet.Add(item)
		}
	}
	for _, item := range st2.GetItems() {
		if !dropSet.In(item) {
			symmetricDifferenceSet.Add(item)
		}
	}
	return symmetricDifferenceSet
}
