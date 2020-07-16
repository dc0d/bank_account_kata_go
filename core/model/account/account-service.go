package account

import (
	"time"
)

type AccountService struct {
	repo AccountRepo
}

func NewAccountService(repo AccountRepo) *AccountService {
	res := &AccountService{
		repo: repo,
	}

	return res
}

func (svc *AccountService) Deposit(id AccountID, amount Amount, at time.Time) error {
	account := NewAccount(id)
	account.Deposit(amount, at)
	return svc.repo.SaveAccount(account)
}

func (svc *AccountService) Withdraw(id AccountID, amount Amount, at time.Time) error {
	account := NewAccount(id)
	account.Withdraw(amount, at)
	return svc.repo.SaveAccount(account)
}

func (svc *AccountService) PrintBankStatement(id AccountID) (Statement, error) {
	acc, err := svc.repo.FindAccount(id)
	if err != nil {
		return Statement{}, err
	}

	return acc.Statement(), nil
}

type AccountServiceInterface interface {
	Deposit(AccountID, Amount, time.Time) error
	Withdraw(AccountID, Amount, time.Time) error
	PrintBankStatement(AccountID) (Statement, error)
}
