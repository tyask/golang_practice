package main

import "testing"

/*
[演習]
IsPrimeメソッドを実装してください。
*/

func IsPrime(n int) bool {
	// TODO nが素数であればtrue、そうでなければfalseを返すように実装してください。
	// なお、素数とは、2以上の自然数で正の約数が1と自分自身のみであるもののことです。
	return false
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
