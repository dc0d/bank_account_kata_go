package actions

import (
	"errors"

	"github.com/dc0d/bank_account_kata_go/core/actions/boundaries"
	"github.com/dc0d/bank_account_kata_go/core/model/account"

	"github.com/dc0d/bank_account_kata_go/test/support"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func DescribePrintBankStatement() {
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
