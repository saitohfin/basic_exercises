package pointers

//Wallet to keep money
type Wallet struct {
	balance int
}

//Deposit add money
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

//Balance returns the current money
func (w Wallet) Balance() int {
	return w.balance
}
