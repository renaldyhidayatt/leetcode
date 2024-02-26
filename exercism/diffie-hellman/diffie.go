package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

func PrivateKey(p *big.Int) *big.Int {
	min := big.NewInt(2)

	max := new(big.Int).Sub(p, min)

	n, err := rand.Int(rand.Reader, max)

	if err != nil {
		panic(err)
	}

	return n.Add(n, min)
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	pp := big.NewInt(g)

	return pp.Exp(pp, private, p)
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {

	private := PrivateKey(p)

	public := PublicKey(private, p, g)

	return private, public

}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	s := new(big.Int)

	return s.Exp(public2, private1, p)
}
