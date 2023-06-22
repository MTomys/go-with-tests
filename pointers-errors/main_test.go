package pointerserrors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, expected Bitcoin) {
		t.Helper()
		received := wallet.Balance()
		if expected != received {
			t.Errorf("expected: %s, received: %s", expected, received)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		expected := Bitcoin(10)
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, expected)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		expected := Bitcoin(10)

		assertBalance(t, wallet, expected)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		if err == nil {
			t.Error("Expected an error when withdrawing more than is deposited")
		}

		assertBalance(t, wallet, startingBalance)
	})
}
