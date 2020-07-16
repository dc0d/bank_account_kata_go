package account

import (
	"fmt"
	"strings"
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

//

type Transactions []Transaction

type Transaction struct {
	Type   TransactionType `json:"type"`
	Amount Amount          `json:"amount"`
	Time   time.Time       `json:"time"`
}

type TransactionType string

const (
	DepositTransaction    TransactionType = "DEPOSIT"
	WithdrawalTransaction TransactionType = "WITHDRAWAL"
)

//

type Amount int

func (amount Amount) String() string {
	if amount == 0 {
		return ""
	}
	return fmt.Sprintf("%.2f", float64(amount))
}

//

type Statement struct {
	Lines []StatementLine
}

func (statement *Statement) AddStatementLine(line StatementLine) {
	statement.Lines = append(statement.Lines, line)
}

func (statement Statement) String() string {
	var sb strings.Builder
	sb.WriteString("date || credit || debit || balance\n")
	if len(statement.Lines) > 0 {
		for i := len(statement.Lines) - 1; i >= 0; i = i - 1 {
			sb.WriteString(statement.Lines[i].String() + "\n")
		}
	}
	return strings.TrimRight(sb.String(), "\n")
}

type StatementLine struct {
	Date    time.Time
	Credit  Amount
	Debit   Amount
	Balance Amount
}

func (sl StatementLine) String() string {
	parts := []string{
		sl.Date.Format("02/01/2006"),
		prependSpace(sl.Credit.String()),
		prependSpace(sl.Debit.String()),
		prependSpace(sl.Balance.String()),
	}
	return strings.Join(parts, " ||")
}

func prependSpace(s string) string {
	if s == "" {
		return s
	}
	return " " + s
}
