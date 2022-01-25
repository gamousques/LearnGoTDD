package main

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	var amount Bitcon = 10
	wallet.Deposit(amount)

	got := wallet.Balance()
	var want Bitcon =   10

	if got != want {
		t.Errorf("got: %d want:%d", got, want)
	}
}