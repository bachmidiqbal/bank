package bankcore

import (
	"encoding/json"
	"errors"
)

type Customer struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type Account struct {
	Customer
	Number  int32
	Balance float64
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit should be greater than zero")
	}

	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to withdraw should be greater than zero")
	}

	if amount > a.Balance {
		return errors.New("the amount to withdraw should be less or equal current balance")
	}

	a.Balance -= amount

	return nil
}

func (a *Account) Statement() string {
	res, _ := json.Marshal(a)
	return string(res)
}
