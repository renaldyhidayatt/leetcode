package sqrt

import "math"

type SqrtDecomposition[E any, Q any] struct {
	querySingleElement func(element E) Q
	unionQ             func(q1 Q, q2 Q) Q
	updateQ            func(oldQ Q, oldE E, newE E) (newQ Q)

	elements  []E
	blocks    []Q
	blockSize uint64
}

func NewSqrtDecomposition[E any, Q any](
	elements []E,
	querySingleElement func(element E) Q,
	unionQ func(q1 Q, q2 Q) Q,
	updateQ func(oldQ Q, oldE E, newE E) (newQ Q),
) *SqrtDecomposition[E, Q] {
	sqrtDec := &SqrtDecomposition[E, Q]{
		querySingleElement: querySingleElement,
		unionQ:             unionQ,
		updateQ:            updateQ,
		elements:           elements,
	}
	sqrt := math.Sqrt(float64(len(sqrtDec.elements)))
	blockSize := uint64(sqrt)
	numBlocks := uint64(math.Ceil(float64(len(elements)) / float64(blockSize)))
	sqrtDec.blocks = make([]Q, numBlocks)
	for i := uint64(0); i < uint64(len(elements)); i++ {
		if i%blockSize == 0 {
			sqrtDec.blocks[i/blockSize] = sqrtDec.querySingleElement(elements[i])
		} else {
			sqrtDec.blocks[i/blockSize] = sqrtDec.unionQ(sqrtDec.blocks[i/blockSize], sqrtDec.querySingleElement(elements[i]))
		}
	}
	sqrtDec.blockSize = blockSize
	return sqrtDec
}

func (s *SqrtDecomposition[E, Q]) Query(start uint64, end uint64) Q {
	firstIndexNextBlock := ((start / s.blockSize) + 1) * s.blockSize
	q := s.querySingleElement(s.elements[start])
	if firstIndexNextBlock > end {
		start++
		for start < end {
			q = s.unionQ(q, s.querySingleElement(s.elements[start]))
			start++
		}
	} else {
		start++
		for start < firstIndexNextBlock {
			q = s.unionQ(q, s.querySingleElement(s.elements[start]))
			start++
		}

		endBlock := end / s.blockSize
		for i := firstIndexNextBlock / s.blockSize; i < endBlock; i++ {
			q = s.unionQ(q, s.blocks[i])
		}

		for i := endBlock * s.blockSize; i < end; i++ {
			q = s.unionQ(q, s.querySingleElement(s.elements[i]))
		}
	}
	return q
}

func (s *SqrtDecomposition[E, Q]) Update(index uint64, newElement E) {
	i := index / s.blockSize
	s.blocks[i] = s.updateQ(s.blocks[i], s.elements[index], newElement)
	s.elements[index] = newElement
}
