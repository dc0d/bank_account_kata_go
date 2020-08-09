package account

import "time"

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
