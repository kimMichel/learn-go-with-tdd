package pointer

import "fmt"

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(quantity Bitcoin) {
	w.balance += quantity
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(quantity Bitcoin) {
	w.balance -= quantity
}
 
type Stringer interface {
	String() string
}