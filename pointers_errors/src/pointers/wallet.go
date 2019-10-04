package pointers

//Wallet to keep money
type Wallet struct {
	balance Bitcoin
}

//Deposit add money
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

//Withdraw deduct money
func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

//Balance returns the current money
func (w Wallet) Balance() Bitcoin {
	return w.balance
}
