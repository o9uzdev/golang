package main

import (
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func main() {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(100))

	fmt.Println(wallet.Balance())

}