package main

import "testing"

func TestIsProbabilisticPrime(t *testing.T) {
	n := uint64(561)
	k := uint64(10)
	result, err := IsProbabilisticPrime(n, k)
	if err != nil {
		t.Errorf("IsProbabilisticPrime(%d, %d) failed: %s", n, k, err)
	}
	if result {
		t.Errorf("IsProbabilisticPrime(%d, %d) failed: got %t, want %t", n, k, result, false)
	}
}

func BenchmarkIsProbabilisticPrime(b *testing.B) {
	n := uint64(6623150861)
	k := uint64(10)
	for i := 0; i < b.N; i++ {
		_, err := IsProbabilisticPrime(n, k)
		if err != nil {
			return
		}
	}
}

func BenchmarkIsPrime(b *testing.B) {
	n := uint64(6623150861)
	for i := 0; i < b.N; i++ {
		IsPrime(n)
	}
}
