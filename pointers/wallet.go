package main

import "fmt"

type Wallet struct {
	balance int
}

func (w *Wallet) Balance() int {
	// Need to use pointer!
	return w.balance
}

func (w *Wallet) Deposit(amount int) {
	// Need to use pointer so that balance lives at the same memory location
	fmt.Printf("address of balance in Deposit is %v\n", &w.balance)
	w.balance += amount
}
