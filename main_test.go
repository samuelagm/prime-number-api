package main

import "testing"

func TestProducePrime(t *testing.T) {
	primeNumber := ProducePrime()
	if primeNumber == 0 && primeNumber < 2 {
		t.Errorf("Number should be prime")
	}
}
