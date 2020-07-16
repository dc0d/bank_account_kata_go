package actions

import (
	"github.com/dc0d/bank_account_kata_go/core/actions/boundaries"
	"github.com/dc0d/bank_account_kata_go/core/model/account"
)

type Withdrawal struct {
	svc account.AccountServiceInterface
}

func NewWithdrawal(svc account.AccountServiceInterface) *Withdrawal {
	res := &Withdrawal{
		svc: svc,
	}

	return res
}

func (action *Withdrawal) Execute(actionInput interface{}, actionOutput interface{}) {
	input := actionInput.(boundaries.WithdrawalInput)
	output := actionOutput.(boundaries.WithdrawalOutput)

	_, _ = input, output

	output(action.svc.Withdraw(input.AccountID, input.Amount, input.Time))
}
