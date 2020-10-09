package primes

import (
	"reflect"
	"testing"
)

func TestSieveOfEratosthenes(t *testing.T) {
	primes := []int { 2, 3, 5, 7, 11, 13, 17, 19, 23, 29 }
	if !reflect.DeepEqual(SieveOfEratosthenes(30), primes) {
		t.Errorf("Result shoul be %v", primes)
	}
}
