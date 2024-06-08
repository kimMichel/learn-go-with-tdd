package pointer

import (
	"testing"
)

func TestWallet(t *testing.T) {
	confirmBalance := func(t *testing.T, wallet Wallet, expected Bitcoin) {
		t.Helper()
		result := wallet.Balance()

		if result != expected {
			t.Errorf("expected: %s, result: %s", expected, result)
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		expected := Bitcoin(10)

		confirmBalance(t, wallet, expected)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		expected := Bitcoin(10)

		confirmBalance(t, wallet, expected)
	})
}