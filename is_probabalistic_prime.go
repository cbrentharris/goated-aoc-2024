package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func IsProbabilisticPrime(n uint64, k uint64) (bool, error) {
	// Miller-Rabin primality test
	if n < 0 {
		return false, fmt.Errorf("IsProbabilisticPrime: n must be non-negative")
	}
	if n <= 3 {
		return true, nil
	}

	if n == 4 {
		return false, nil
	}
	d := n - 1
	//fmt.Printf("d: %d\n", d)

	for d%2 != 0 {
		d /= 2
	}

	for i := uint64(0); i < k; i++ {
		if !MillerTest(n, d) {
			return false, nil
		}
	}
	return true, nil
}

func MillerTest(n, d uint64) bool {
	a, _ := rand.Int(rand.Reader, new(big.Int).SetUint64(n-4))
	a.Add(a, big.NewInt(2))
	x := ModPow(a.Uint64(), d, n)
	if x == 1 || x == n-1 {
		return true
	}
	for j := d; j < n-1; j *= 2 {
		x = ModPow(x, 2, n)
		if x == 1 {
			return false
		}
		if x == n-1 {
			return true
		}
	}
	return false
}

func ModPow(a, b, m uint64) uint64 {
	result := new(big.Int).Exp(
		new(big.Int).SetUint64(a),
		new(big.Int).SetUint64(b),
		new(big.Int).SetUint64(m),
	)
	return result.Uint64()
}

func IsPrime(n uint64) bool {
	if n < 0 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := uint64(3); i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
