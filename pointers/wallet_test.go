package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Euro) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Euro(10))
		assertBalance(t, wallet, Euro(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Euro(20)}
		wallet.Withdraw(Euro(10))
		assertBalance(t, wallet, Euro(10))
	})
}
