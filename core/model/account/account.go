package account

import (
	"time"
)

type Account struct {
	AccountID    AccountID    `json:"account_id"`
	Transactions Transactions `json:"transactions,omitempty"`
}

func NewAccount(id AccountID) *Account {
	res := &Account{
		AccountID: id,
	}

	return res
}

func (account *Account) Deposit(amount Amount, at time.Time) {
	transaction := Transaction{
		Type:   DepositTransaction,
		Amount: amount,
		Time:   at,
	}

	account.Transactions = append(account.Transactions, transaction)
}

func (account *Account) Withdraw(amount Amount, at time.Time) {
	transaction := Transaction{
		Type:   WithdrawalTransaction,
		Amount: amount,
		Time:   at,
	}

	account.Transactions = append(account.Transactions, transaction)
}

func (account *Account) Statement() Statement {
	var st Statement

	var balance Amount
	for _, tx := range account.Transactions {
		var line StatementLine
		line.Date = tx.Time

		switch tx.Type {
		case DepositTransaction:
			line.Credit = tx.Amount
			balance += tx.Amount
		case WithdrawalTransaction:
			line.Debit = tx.Amount
			balance -= tx.Amount
		}
		line.Balance = balance

		st.AddStatementLine(line)
	}
	return st
}

type (
	AccountID string
)
