package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Euro(10))
		assertBalance(t, wallet, Euro(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Euro(20)}
		err := wallet.Withdraw(Euro(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Euro(10))
	})

	t.Run("overdraw", func(t *testing.T) {
		startingBalance := Euro(10) // Note that as we are using this in assertBalance, we need to use Euro
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Euro(20))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Euro) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, err, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("wanted an error, got none")
	}
	if err != want {
		t.Errorf("got %q, want %q", err, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}
