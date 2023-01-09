package util

import (
	crypto "crypto/rand"
	"math/big"
)

func NewCryptoRand(min, max int) int64 {
	num, err := crypto.Int(crypto.Reader, big.NewInt(int64(max)-int64(min)))
	if err != nil {
		panic(err)
	}
	return num.Int64()
}
