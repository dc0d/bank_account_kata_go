package actions

import (
	"github.com/dc0d/bank_account_kata_go/core/actions/boundaries"
	"github.com/dc0d/bank_account_kata_go/core/model/account"
)

type Deposit struct {
	svc account.AccountServiceInterface
}

func NewDeposit(svc account.AccountServiceInterface) *Deposit {
	res := &Deposit{
		svc: svc,
	}

	return res
}

func (action *Deposit) Execute(actionInput interface{}, actionOutput interface{}) {
	input := actionInput.(boundaries.DepositInput)
	output := actionOutput.(boundaries.DepositOutput)

	output(action.svc.Deposit(input.AccountID, input.Amount, input.Time))
}
