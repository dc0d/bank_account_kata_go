//go:generate moq -pkg actions -out ./account_service_spy_test.go ./../model/account AccountServiceInterface:AccountServiceInterfaceSpy
//go:generate moq -pkg actions -out ./account_repo_spy_test.go ./../model/account AccountRepo:AccountRepoSpy

// a black & white test
// AccountServiceInterfaceSpy is a spy that records call inputs and outputs
// so it can be used as a mock to veryfy interaction
// this way the whitebox tests are done for designing
// also AccountServiceInterfaceSpy can call the actual AccountService and acts as a proxy
// (maybe from the testing point of view, it should be considered a decorator)
// this way the blackbox test is done

// don't do this, all this repo is just practicing and playing around with ideas
// usually whitebix and blackbox testing is about testing one thing, not cross module/package
// but I wanted to see how things will turn out

package actions

import (
	"errors"
	"testing"
	"time"

	"github.com/dc0d/bank_account_kata_go/core/actions/boundaries"
	"github.com/dc0d/bank_account_kata_go/core/model/account"

	"github.com/dc0d/bank_account_kata_go/test/support"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test_actions_suite_unit(t *testing.T) {
	ActionsSuite()

	RegisterFailHandler(Fail)
	RunSpecs(t, "Actions Suite")
}

func ActionsSuite() {
	Describe("Deposit", func() {
		var (
			depositAction *Deposit

			realAccountService *account.AccountService
			accountService     *AccountServiceInterfaceSpy

			input  boundaries.DepositInput
			output boundaries.DepositOutput

			resultedError error
		)

		BeforeEach(func() {
			accountRepo := &AccountRepoSpy{}
			accountRepo.SaveAccountFunc = func(in1 *account.Account) error {
				return nil
			}
			realAccountService = account.NewAccountService(accountRepo)

			accountService = &AccountServiceInterfaceSpy{}
			accountService.DepositFunc = func(in1 account.AccountID, in2 account.Amount, in3 time.Time) error {
				return realAccountService.Deposit(in1, in2, in3)
			}

			depositAction = NewDeposit(accountService)
		})

		Context("When the deposit action is call", func() {
			BeforeEach(func() {
				input = boundaries.DepositInput{
					AccountID: "ACCOUNT_ID_2",
					Amount:    500,
					Time:      support.ParseDate("10-01-2012"),
				}

				output = func(err error) { resultedError = err }
			})

			BeforeEach(func() {
				depositAction.Execute(input, output)
			})

			It("Should deposit the amount in bank account", func() {
				Expect(resultedError).NotTo(HaveOccurred())

				callDepositFuncOnce := len(accountService.DepositCalls()) == 1
				Expect(callDepositFuncOnce).To(BeTrue(), `Deposit func should be called once`)
				Expect(accountService.DepositCalls()[0].In1).To(Equal(input.AccountID))
				Expect(accountService.DepositCalls()[0].In2).To(Equal(input.Amount))
				Expect(accountService.DepositCalls()[0].In3).To(Equal(input.Time))
			})
		})
	})

	Describe("Withdrawal", func() {
		var (
			withdrawalAction *Withdrawal

			realAccountService *account.AccountService
			accountService     *AccountServiceInterfaceSpy

			input  boundaries.WithdrawalInput
			output boundaries.WithdrawalOutput

			resultedError = errors.New("SHOULD GO AWAY")
		)

		BeforeEach(func() {
			accountRepo := &AccountRepoSpy{}
			accountRepo.SaveAccountFunc = func(in1 *account.Account) error {
				return nil
			}
			realAccountService = account.NewAccountService(accountRepo)

			accountService = &AccountServiceInterfaceSpy{}
			accountService.WithdrawFunc = func(in1 account.AccountID, in2 account.Amount, in3 time.Time) error {
				return realAccountService.Withdraw(in1, in2, in3)
			}

			withdrawalAction = NewWithdrawal(accountService)
		})

		Context("When the withdrawal action is call", func() {
			BeforeEach(func() {
				input = boundaries.WithdrawalInput{
					AccountID: "ACCOUNT_ID_2",
					Amount:    500,
					Time:      support.ParseDate("10-01-2012"),
				}

				output = func(err error) { resultedError = err }
			})

			BeforeEach(func() {
				withdrawalAction.Execute(input, output)
			})

			It("Should withdrawal the amount in bank account", func() {
				Expect(resultedError).NotTo(HaveOccurred())

				callwithdrawalFuncOnce := len(accountService.WithdrawCalls()) == 1
				Expect(callwithdrawalFuncOnce).To(BeTrue(), `withdrawal func should be called once`)
				Expect(accountService.WithdrawCalls()[0].In1).To(Equal(input.AccountID))
				Expect(accountService.WithdrawCalls()[0].In2).To(Equal(input.Amount))
				Expect(accountService.WithdrawCalls()[0].In3).To(Equal(input.Time))
			})
		})
	})

	Describe("Print Bank Statement", func() {
		var (
			action *PrintBankStatement

			realAccountService *account.AccountService
			accountService     *AccountServiceInterfaceSpy

			input  boundaries.PrintBankStatementInput
			output boundaries.PrintBankStatementOutput

			resultedError     = errors.New("SHOULD GO AWAY")
			resultedStatement account.Statement
		)

		BeforeEach(func() {
			accountRepo := &AccountRepoSpy{}
			accountRepo.FindAccountFunc = func(in1 account.AccountID) (*account.Account, error) {
				acc := account.NewAccount(in1)
				acc.Deposit(1000, support.ParseDate("10-01-2012"))
				acc.Deposit(2000, support.ParseDate("13-01-2012"))
				acc.Withdraw(500, support.ParseDate("14-01-2012"))

				return acc, nil
			}

			realAccountService = account.NewAccountService(accountRepo)

			accountService = &AccountServiceInterfaceSpy{}
			accountService.PrintBankStatementFunc = func(in1 account.AccountID) (account.Statement, error) {
				return realAccountService.PrintBankStatement(in1)
			}

			action = NewPrintBankStatement(accountService)
		})

		Context("When PrintBankStatement is called", func() {
			BeforeEach(func() {
				input = "ACCOUNT_1"

				output = func(statement account.Statement, err error) {
					resultedStatement = statement
					resultedError = err
				}
			})

			BeforeEach(func() {
				action.Execute(input, output)
			})

			It("Should give the expected statement", func() {
				Expect(resultedError).NotTo(HaveOccurred())

				callPrintBankStatementOnce := len(accountService.PrintBankStatementCalls()) == 1
				Expect(callPrintBankStatementOnce).To(BeTrue(), `PrintBankStatement func should be called once`)

				var expectedStatement account.Statement
				for _, line := range sampleStatementLines() {
					expectedStatement.AddStatementLine(line)
				}

				Expect(resultedStatement).To(Equal(expectedStatement))
			})
		})
	})
}

func sampleStatementLines() []account.StatementLine {
	return []account.StatementLine{
		{Date: support.ParseDate("10-01-2012"), Credit: 1000, Debit: 0, Balance: 1000},
		{Date: support.ParseDate("13-01-2012"), Credit: 2000, Debit: 0, Balance: 3000},
		{Date: support.ParseDate("14-01-2012"), Credit: 0, Debit: 500, Balance: 2500},
	}
}
