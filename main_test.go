package bank_test

import (
	"example.com/bank"
	"testing"
)

func TestDeposit(t *testing.T) {
	acc := bank.NewAccount()
	acc.Deposit(100)
	want := 100
	got := acc.Balance

	assertBalance(t, want, got)
}

func TestCantWithdrawNegative(t *testing.T) {
	acc := bank.NewAccount()
	err := acc.Withdraw(100)

	if err == nil {
		t.Errorf("want error, got nil")
	}

	assertBalance(t, 0, acc.Balance)
}

func TestWithdraw(t *testing.T) {
	acc := bank.NewAccount()
	acc.Deposit(200)
	err := acc.Withdraw(100)

	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	assertBalance(t, 100, acc.Balance)
}

func TestTransfer(t *testing.T) {
	accFrom := bank.NewAccount()
	accTo := bank.NewAccount()

	accFrom.Deposit(350)
	err := accFrom.Transfer(&accTo, 100)

	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	assertBalance(t, 250, accFrom.Balance)
	assertBalance(t, 100, accTo.Balance)
}

func assertBalance(t testing.TB, want, got int) {
	t.Helper()

	if got != want {
		t.Errorf("want balance %v, got %v", want, got)
	}
}
