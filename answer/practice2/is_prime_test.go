package main

import "testing"

/*
[演習]
IsPrimeメソッドを実装してください。
*/

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func TestIsPrime(t *testing.T) {
	if IsPrime(0) {
		t.Error("0 is not a prime number")
	}
	if IsPrime(1) {
		t.Error("1 is not a prime number")
	}
	if !IsPrime(2) {
		t.Error("2 is a prime number")
	}
	if !IsPrime(13) {
		t.Error("13 is a prime number")
	}
	if IsPrime(91) {
		t.Error("91 is not a prime number")
	}
	if IsPrime(1000007) {
		t.Error("1000007 is not prime number")
	}
}
