package bank

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
)

type Account struct {
	accountId int32
	balance   decimal.Decimal
}

// canWidthdraw check is sufficient money is in the account
func (a *Account) canWithdraw(amount decimal.Decimal) (bool) {
	return a.balance.Cmp(amount) >= 0
}

// isPositiveDecimal checks if the decimal is positive
func isNegativeDecimal(amount decimal.Decimal) (bool) {
	return amount.Cmp(decimal.NewFromFloat(0)) < 0
}

func (a *Account) Withdraw(amount decimal.Decimal) (decimal.Decimal, error) {
	if isNegativeDecimal(amount) {
		return a.balance, errors.New(fmt.Sprintf("Unable to withdraw values less than or equal 0"))
	}

	if !a.canWithdraw(amount) {
		return a.balance, errors.New(fmt.Sprintf("Unable to withdraw more than the balance"))
	}

	a.balance = a.balance.Sub(amount)

	return a.balance, nil
}

func (a *Account) Deposit(amount decimal.Decimal) (decimal.Decimal, error) {
	if isNegativeDecimal(amount) {
		return a.balance, errors.New(fmt.Sprintf("Unable to deposit values less than or equal 0"))
	}

	a.balance = a.balance.Add(amount)
	return a.balance, nil
}

func (a Account) GetBalance() decimal.Decimal {
	return a.balance
}

func CreateAccount(id int32) Account {
	return Account{accountId: id, balance: decimal.NewFromFloat(0.0)}
}
