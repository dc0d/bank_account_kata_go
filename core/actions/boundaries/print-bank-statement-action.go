package boundaries

import "github.com/dc0d/bank_account_kata_go/core/model/account"

type (
	PrintBankStatementInput account.AccountID

	PrintBankStatementOutput func(account.Statement, error)
)
