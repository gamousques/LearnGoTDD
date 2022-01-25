package main

import "testing"

func TestWallet(t *testing.T) {
	
	t.Run("Deposit amount balance increased by amount", func(t *testing.T) {
		wallet := Wallet{}
		amount := Bitcoin(10)
		wallet.Deposit(amount)
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
		
	})

	t.Run("Withdraw amount balance decreaced by amount", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(10))
		
		want := Bitcoin(0)
		assertNoError(t, err)
		assertBalance(t, wallet, want)

		
	})

	t.Run("Witdraw so balance throws error because of insufficient funds and does not decrease balance", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(30))
		assertError(t, err, ErrInsufficientFunds.Error())
		assertBalance(t,wallet, startingBalance)
	
	})
}

func assertBalance (t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("got: %s want:%s", got, want)
	}
}

func assertError(t testing.TB, got error, want string)  {
	t.Helper()
	if got == nil {
		t.Fatal("didnt get an error but was expecting one!")
	}
	if got.Error() != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}

func assertNoError(t testing.TB, got error)  {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but did not want one!")
	}
}

