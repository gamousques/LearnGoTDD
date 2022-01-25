package main

type Bitcon int

type Wallet struct {
	balance Bitcon
}

func (w *Wallet) Deposit(amount Bitcon) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcon {
	return w.balance
}