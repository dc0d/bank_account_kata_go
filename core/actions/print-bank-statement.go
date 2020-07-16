package actions

import (
	"github.com/dc0d/bank_account_kata_go/core/actions/boundaries"
	"github.com/dc0d/bank_account_kata_go/core/model/account"
)

type PrintBankStatement struct {
	svc account.AccountServiceInterface
}

func NewPrintBankStatement(svc account.AccountServiceInterface) *PrintBankStatement {
	res := &PrintBankStatement{
		svc: svc,
	}

	return res
}

func (action *PrintBankStatement) Execute(actionInput interface{}, actionOutput interface{}) {
	input := actionInput.(boundaries.PrintBankStatementInput)
	output := actionOutput.(boundaries.PrintBankStatementOutput)

	output(action.svc.PrintBankStatement(account.AccountID(input)))
}
