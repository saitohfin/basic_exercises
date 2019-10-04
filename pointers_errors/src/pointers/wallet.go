package pointers

import "errors"

//ErrInsufficientFunds when you haven't enough balance
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

//Wallet to keep money
type Wallet struct {
	balance Bitcoin
}

//Deposit add money
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

//Withdraw deduct money
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

//Balance returns the current money
func (w Wallet) Balance() Bitcoin {
	return w.balance
}
