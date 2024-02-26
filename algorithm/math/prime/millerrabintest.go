package prime

import (
	"algoritmAndDs/math/modular"
	"math/rand"
)

func formatNum(num int64) (d int64, s int64) {
	d = num - 1

	for d%2 == 0 {
		d /= 2
		s++
	}
	return
}

func isTrival(num int64) (prime bool, trivial bool) {
	if num <= 4 {
		prime = num == 2 || num == 3
		trivial = true
	} else {
		prime = false
		trivial = num%2 == 0
	}

	return
}

func MillerTest(num, witness int64) (bool, error) {
	d, _ := formatNum(num)

	res, err := modular.Exponentiation(witness, d, num)

	if err != nil {
		return false, err
	}

	if res == 1 || res == num-1 {
		return true, nil
	}

	for d != num-1 {
		res = (res * res) % num
		d *= 2

		if res == 1 {
			return false, nil
		}

		if res == num-1 {
			return true, nil
		}
	}

	return false, nil
}

func MillerRandomTest(num int64) (bool, error) {
	random := rand.Int63n(num-2) + 2

	return MillerTest(num, random)
}

func MillerTestMultiple(num int64, witnesses ...int64) (bool, error) {
	for _, witness := range witnesses {
		prime, err := MillerTest(num, witness)

		if err != nil {
			return false, err
		}

		if !prime {
			return false, nil
		}
	}

	return true, nil
}

func MillerRabitProbalistic(num, rounds int64) (bool, error) {
	if prime, trival := isTrival(num); trival {
		return prime, nil
	}

	for i := int64(0); i < rounds; i++ {
		val, err := MillerRandomTest(num)

		if err != nil {
			return false, err
		}

		if !val {
			return false, nil
		}
	}

	return true, nil
}

func MillerRabinDeterministic(num int64) (bool, error) {
	if prime, trival := isTrival(num); trival {
		return prime, nil
	}

	switch {
	case num < 2047:
		return MillerTest(num, 2)
	case num < 1_373_653:
		return MillerTestMultiple(num, 2, 3)
	case num < 9_080_191:
		return MillerTestMultiple(num, 31, 73)
	case num < 25_326_001:
		return MillerTestMultiple(num, 2, 3, 5)
	case num < 1_122_004_669_633:
		return MillerTestMultiple(num, 2, 13, 23, 1_662_803)
	case num < 2_152_302_898_747:
		return MillerTestMultiple(num, 2, 3, 5, 7, 11)
	case num < 341_550_071_728_321:
		return MillerTestMultiple(num, 2, 3, 5, 7, 11, 13, 17)
	default:
		return MillerTestMultiple(num, 2, 3, 5, 7, 11, 13, 17, 19, 23, 31, 37)
	}
}
