package day22

import (
	"goated-aoc-2024/year2024"
	"testing"
)

var exampleInput = `1
10
100
2024`

var exampleInput2 = `1
2
3
2024`

func TestCalculateNthSecretExample(t *testing.T) {
	result := CalculateNextSecret(123)
	if result != 15887950 {
		t.Errorf("Expexted 15887950 isntead found %d", result)
	}
}

func TestSumOfNthSecretsExample(t *testing.T) {
	result := SumOfNthSecrets(exampleInput, 2000)
	if result != 37327623 {
		t.Errorf("Expexted 37327623 isntead found %d", result)
	}
}

func TestMaxBananasExample(t *testing.T) {
	result := MaxBananas(exampleInput2, 2000)
	if result != 23 {
		t.Errorf("Expected 23 bananas, instead found %d", result)
	}
}

func TestMaxBananas(t *testing.T) {
	result := MaxBananas(year2024.ReadInput("input.txt"), 2000)
	if result <= 1678 {
		t.Errorf("Expected more than 1678 bananas, instead found %d", result)
	}
	if result != 1710 {
		t.Errorf("Expected 1710 bananas, instead found %d", result)
	}
}

func TestMaxBananasExample2(t *testing.T) {
	result := MaxBananas("123", 10)
	if result != 6 {
		t.Errorf("Expected 6 bananas, instead found %d", result)
	}
}

func TestSumOfNthSecrets(t *testing.T) {
	result := SumOfNthSecrets(year2024.ReadInput("input.txt"), 2000)

	if result != 15006633487 {
		t.Errorf("Expexted 15006633487 isntead found %d", result)
	}
}
