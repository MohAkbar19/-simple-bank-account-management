package bank

import "errors"

// Account represents a bank account with a balance.
type Account struct {
    Balance float64
}

// Deposit adds an amount to the account balance.
// Returns an error if the amount is zero or negative.
func (a *Account) Deposit(amount float64) error {
    if amount <= 0 {
        return errors.New("jumlah deposit harus positif")
    }
    a.Balance += amount
    return nil
}

// Withdraw subtracts an amount from the account balance.
// Returns an error if the amount is zero, negative, or exceeds the balance.
func (a *Account) Withdraw(amount float64) error {
    if amount <= 0 {
        return errors.New("jumlah penarikan harus positif")
    }
    if amount > a.Balance {
        return errors.New("saldo tidak mencukupi")
    }
    a.Balance -= amount
    return nil
}

// GetBalance returns the current balance of the account.
func (a *Account) GetBalance() float64 {
    return a.Balance
}
