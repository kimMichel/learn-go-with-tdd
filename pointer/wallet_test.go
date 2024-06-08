package pointer

import (
	"testing"
)

func TestWallet(t *testing.T) {
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

	t.Run("Withdraw with not enought balance", func(t *testing.T) {
		initialBalance := Bitcoin(20)
		wallet := Wallet{initialBalance}
		err := wallet.Withdraw(100)

		confirmBalance(t, wallet, initialBalance)
		confirmError(t, err, ErrorNotEnoughBalance)
	})
}

func confirmBalance(t *testing.T, wallet Wallet, expected Bitcoin) {
	t.Helper()
	result := wallet.Balance()

	if result != expected {
		t.Errorf("expected: %s, result: %s", expected, result)
	}
}

func confirmError(t *testing.T, result error, expected error)  {
	t.Helper()
	if result == nil {
		t.Fatal("Expected an error..")
	}

	if result != expected {
		t.Errorf("expected: %s, result: %s", expected, result)
	}
}