package webapi

import (
	"time"

	"github.com/dc0d/bank_account_kata_go/core/actions/boundaries"
	"github.com/dc0d/bank_account_kata_go/core/model/account"
)

type DepositActionProvider interface {
	ProvideDepositAction() boundaries.Action
}

type DepositCommandPayload struct {
	CommandPayload
	AccountID account.AccountID `json:"account_id"`
	Amount    account.Amount    `json:"amount"`
	Time      time.Time         `json:"time"`
}

func (cmd *DepositCommandPayload) Init() {
	cmd.Command = DepositCommandTerm
}

//

type WithdrawalActionProvider interface {
	ProvideWithdrawalAction() boundaries.Action
}

type WithdrawCommandPayload struct {
	CommandPayload
	AccountID account.AccountID `json:"account_id"`
	Amount    account.Amount    `json:"amount"`
	Time      time.Time         `json:"time"`
}

func (cmd *WithdrawCommandPayload) Init() {
	cmd.Command = WithdrawCommandTerm
}

//

type PrintBankStatementActionProvider interface {
	ProvidePrintBankStatementAction() boundaries.Action
}

//

type (
	CommandPayload struct {
		Command CommandTerm `json:"command"`
	}

	CommandTerm string
)

const (
	DepositCommandTerm  CommandTerm = "DEPOSIT_AMOUNT"
	WithdrawCommandTerm CommandTerm = "WITHDRAW_AMOUNT"
)
