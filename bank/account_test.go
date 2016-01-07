package bank

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeposit(t *testing.T) {

	account := CreateAccount(10)
	assert.Equal(t, account.GetBalance(), decimal.NewFromFloat(0))

	account.Deposit(decimal.NewFromFloat(2510.50))
	assert.Equal(t, account.GetBalance(), decimal.NewFromFloat(2510.50))

	// Attempt to deposit negative value
	_, err := account.Deposit(decimal.NewFromFloat(-1))
	assert.NotNil(t, err)

}

func TestWithdraw(t *testing.T) {

	account := CreateAccount(10)
	assert.Equal(t, account.GetBalance(), decimal.NewFromFloat(0))

	account.Deposit(decimal.NewFromFloat(2510.50))
	account.Withdraw(decimal.NewFromFloat(2510))
	assert.Equal(t, account.GetBalance(), decimal.NewFromFloat(0.50))

	// Attempt to withdraw more than balance
	account.Withdraw(decimal.NewFromFloat(1))
	assert.Equal(t, account.GetBalance(), decimal.NewFromFloat(0.50))

	// Attempt to withdraw negative amount
	_, err := account.Withdraw(decimal.NewFromFloat(-1))
	assert.NotNil(t, err)

}

func TestCanWidthdraw(t *testing.T) {

	account := CreateAccount(10)
	account.Deposit(decimal.NewFromFloat(20))

	account.Withdraw(decimal.NewFromFloat(50))
	assert.Equal(t, decimal.NewFromFloat(20).String(), account.GetBalance().String())

}