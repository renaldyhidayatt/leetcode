package dominoes

type Domino [2]int
type Dmap map[Domino]int

func MakeChain(input []Domino) (rez []Domino, err bool) {
	leftDomino := make(Dmap)

	for _, dmn := range input {
		leftDomino[dmn]++
	}

	for i, dmn := range input {
		if i == 0 {
			rez = append(rez, dmn)

			leftDomino.reduce(dmn)
			continue
		}
		rez = append(rez, leftDomino.find(rez[len(rez)-1][1]))
	}

	if len(leftDomino) > 0 {
		return nil, false
	}

	if len(rez) > 0 {
		if rez[0][0] != rez[len(rez)-1][1] {
			return nil, false
		}
	}

	return rez, true
}

func (dmap Dmap) reduce(dm Domino) {
	if dmap[dm] > 1 {
		dmap[dm]--
	} else {
		delete(dmap, dm)
	}
}

func (dmap Dmap) find(n int) Domino {
	var candidate, candidateOrig Domino

	var candidateOcur int

	for tstDm, val := range dmap {
		if tstDm[0] == n {
			if tstDm[0] == tstDm[1] {
				candidate = tstDm
				candidateOrig = candidate
				break
			}

			if val > candidateOcur {
				candidate = tstDm
				candidateOrig = candidate
				candidateOcur = val
			}
		}

		if tstDm[1] == n {
			if val > candidateOcur {
				candidateOrig = tstDm
				candidate = Domino{
					tstDm[1],
					tstDm[0],
				}
				candidateOcur = val
			}
		}

	}

	dmap.reduce(candidateOrig)

	return candidate

}
