package main

import (
	"testing"
	"time"
)

func TestProducePrime(t *testing.T) {
	primeNumber := ProducePrime()
	if primeNumber == 0 && primeNumber < 2 {
		t.Errorf("Number should be prime")
	}
}

func TestPrimeGenerator(t *testing.T) {
	primeNumbers := generate(10000000, time.Second*2)
	if len(primeNumbers) == 0 {
		t.Errorf("Prime number array should not be empty")
	}
}
