package matrix

func (m1 Matrix[T]) CheckEqual(m2 Matrix[T]) bool {
	if !m1.MatchDimensions(m2) {
		return false
	}

	c := make(chan bool)

	for i := range m1.elements {
		go func(i int) {
			for j := range m1.elements[i] {
				if m1.elements[i][j] != m2.elements[i][j] {
					c <- false
					return
				}
			}
			c <- true
		}(i)
	}

	for range m1.elements {
		if !<-c {
			return false
		}
	}

	return true
}
