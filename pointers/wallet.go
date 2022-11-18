package main

import "fmt"

type Euro int

type Wallet struct {
	balance Euro
}

type Stringer interface {
	String() string
}

func (e Euro) String() string {
	return fmt.Sprintf("%d EUR", e)
}

func (w *Wallet) Balance() Euro {
	// Need to use pointer!
	return w.balance
}

func (w *Wallet) Deposit(amount Euro) {
	// Need to use pointer so that balance lives at the same memory location
	// fmt.Printf("address of balance in Deposit is %v\n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Euro) {
	w.balance -= amount
}
