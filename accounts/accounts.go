package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string // 이것도 첫글자가 소문자면 private
	balance int
}

var errBalance = errors.New("check your balance")

// NewAccount create account
func NewAccount(owner string) *Account {
	account := Account{
		owner:   owner,
		balance: 0,
	}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount on your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errBalance
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// GetOwner of the account
func (a Account) GetOwner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.GetOwner(), "'s account.\nHas: ", a.balance)
}
