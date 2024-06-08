package pointer

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		result := wallet.Balance()
		expected := Bitcoin(10)

		if result != expected {
			t.Errorf("expected: %s, result: %s", expected, result)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		result := wallet.Balance()
		expected := Bitcoin(10)

		if result != expected {
			t.Errorf("expected: %s, result: %s", expected, result)
		}
	})
}