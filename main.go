package bank

import "errors"

type Account struct {
	Balance int
}

func NewAccount() Account {
	return Account{}
}

func (a *Account) Deposit(amount int) {
	a.Balance += amount
}

func (a *Account) Withdraw(amount int) error {
	if a.Balance-amount < 0 {
		return errors.New("can't withdraw when balance is 0")
	}

	a.Balance -= amount

	return nil
}

func (a *Account) Transfer(to *Account, amount int) error {
	err := a.Withdraw(amount)

	if err != nil {
		return err
	}

	to.Balance += amount

	return nil
}
