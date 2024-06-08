package pointer

import (
	"errors"
	"fmt"
)

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

var ErrorNotEnoughBalance = errors.New("its not possible to withdraw: not enought balance")

func (w *Wallet) Withdraw(quantity Bitcoin) error {
	if quantity > w.balance {
		return ErrorNotEnoughBalance
	}

	w.balance -= quantity
	return nil
}
 
type Stringer interface {
	String() string
}