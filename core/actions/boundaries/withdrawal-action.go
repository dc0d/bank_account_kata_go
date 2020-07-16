package boundaries

import (
	"time"

	"github.com/dc0d/bank_account_kata_go/core/model/account"
)

type (
	WithdrawalInput struct {
		AccountID account.AccountID
		Amount    account.Amount
		Time      time.Time
	}

	WithdrawalOutput func(error)
)
