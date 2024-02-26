package math

import (
	"errors"
	"math/big"
)

func DefaultPolynomial(n *big.Int) func(*big.Int) *big.Int {
	bigOne := big.NewInt(1)
	bigTwo := big.NewInt(2)
	return func(x *big.Int) *big.Int {
		xSquared := new(big.Int).Exp(x, bigTwo, n)
		xSquared.Add(xSquared, bigOne)
		xSquared.Mod(xSquared, n)
		return xSquared
	}
}

func PollardsRhoFactorization(n *big.Int, f func(n *big.Int) func(x *big.Int) *big.Int) (*big.Int, error) {
	x, y, d := big.NewInt(2), big.NewInt(2), big.NewInt(1)
	bigOne := big.NewInt(1)
	g := f(n)
	for d.Cmp(bigOne) == 0 {
		x = g(x)
		y = g(g(y))
		sub := new(big.Int).Sub(x, y)
		d.GCD(nil, nil, sub.Abs(sub), n)
	}
	if d.Cmp(n) == 0 {
		return nil, errors.New("factorization failed")
	}
	return d, nil
}
