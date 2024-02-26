package queenattack

import "errors"

func CanQueenAttack(s1, s2 string) (bool, error) {
	if s1 == s2 {
		return false, errors.New("same square")
	}

	s1x := s1[0]
	s1y := s1[1]
	s2x := s2[0]
	s2y := s2[1]

	if s1x < 'a' || s1x > 'h' ||
		s2x < 'a' || s2x > 'h' ||
		s1y < '1' || s1y > '8' ||
		s2y < '1' || s2y > '8' {
		return false, errors.New("out of bounds")
	}

	if s1x == s2x || s1y == s2y {
		return true, nil
	}

	for x := -1; x <= 1; x += 2 {
		for y := -1; y <= 1; y += 2 {
			curX := int(s1x)
			curY := int(s1y)

			for curX >= 'a' && curX <= 'h' && curY >= '1' && curY <= '8' {
				if curX == int(s2x) && curY == int(s2y) {
					return true, nil
				}

				curX += x
				curY += y
			}
		}
	}
	return false, nil
}
