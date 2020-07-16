//go:generate moq -out ./account_repo_spy_test.go ./ AccountRepo:AccountRepoSpy

package account

import (
	"fmt"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/dc0d/bank_account_kata_go/test/support"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func Test_account_suite_unit(t *testing.T) {
	AccountSuite()

	RegisterFailHandler(Fail)
	RunSpecs(t, "Account Suite")
}

func AccountSuite() {
	Describe("Amount", func() {
		DescribeTable("check all room types",
			func(amount Amount, expectedString string) {
				actualString := fmt.Sprint(amount)

				Expect(actualString).To(Equal(expectedString))
			},
			Entry("500", Amount(500), "500.00"),
			Entry("0", Amount(0), ""),
		)
	})

	Describe("Statement", func() {
		DescribeTable("StatementLine to string",
			func(line StatementLine, expectedString string) {
				Expect(line.String()).To(Equal(expectedString))
			},
			Entry("1", sampleStatementLines()[0], expectedStatementLineStrings()[0]),
			Entry("2", sampleStatementLines()[1], expectedStatementLineStrings()[1]),
			Entry("3", sampleStatementLines()[2], expectedStatementLineStrings()[2]),
		)

		var (
			statement Statement
		)

		Context("Statement text output", func() {
			It("Should match the expected statement", func() {
				for _, line := range sampleStatementLines() {
					statement.AddStatementLine(line)
				}

				fromNewest := expectedStatementLineStrings()
				sort.Sort(sort.Reverse(sort.StringSlice(fromNewest)))

				lines := []string{"date || credit || debit || balance"}
				lines = append(lines, fromNewest...)
				expectedStatement := strings.Join(lines, "\n")

				Expect(statement.String()).To(Equal(expectedStatement))
			})
		})
	})

	Describe("Account", func() {
		var (
			id AccountID

			account *Account
		)

		BeforeEach(func() {
			id = "AN_ACCOUNT"

			account = NewAccount(id)
		})

		Context("When deposit an amount", func() {
			var (
				amount              Amount
				at                  time.Time
				expectedTransaction Transaction
			)

			BeforeEach(func() {
				amount = 500
				at = time.Date(2020, 07, 18, 0, 0, 0, 0, time.Local)

				expectedTransaction.Type = DepositTransaction
				expectedTransaction.Amount = amount
				expectedTransaction.Time = at
			})

			BeforeEach(func() {

				account.Deposit(amount, at)
			})

			It("Should show up in account transactions", func() {
				Expect(account.Transactions).To(ContainElement(expectedTransaction))
			})
		})
	})

	Describe("AccountService", func() {
		var (
			id     AccountID
			amount Amount
			at     time.Time

			resultedError error
		)

		var (
			accountRepo    *AccountRepoSpy
			accountService *AccountService
		)

		Context("When deposit an amount", func() {
			BeforeEach(func() {
				id = "ACCOUNT_ID"
				amount = 500
				at = time.Date(2020, 07, 18, 0, 0, 0, 0, time.Local)

				accountRepo = &AccountRepoSpy{}
				accountRepo.SaveAccountFunc = func(in1 *Account) error {
					return nil
				}

				accountService = NewAccountService(accountRepo)
			})

			BeforeEach(func() {
				resultedError = accountService.Deposit(id, amount, at)
			})

			It("Should show up in account repo", func() {
				Expect(resultedError).NotTo(HaveOccurred())

				saveCalledOnce := len(accountRepo.SaveAccountCalls()) == 1
				Expect(saveCalledOnce).To(BeTrue(), `Save func of repo should be called once`)

				expectedAccount := NewAccount(id)
				expectedAccount.Deposit(amount, at)
				Expect(accountRepo.SaveAccountCalls()[0].In1).To(Equal(expectedAccount))

				Expect(len(expectedAccount.Transactions) > 0).To(BeTrue(), "no transaction added")
			})
		})

		Context("When withdraw an amount", func() {
			BeforeEach(func() {
				id = "ACCOUNT_ID"
				amount = 500
				at = time.Date(2020, 07, 18, 0, 0, 0, 0, time.Local)

				accountRepo = &AccountRepoSpy{}
				accountRepo.SaveAccountFunc = func(in1 *Account) error {
					return nil
				}

				accountService = NewAccountService(accountRepo)
			})

			BeforeEach(func() {
				resultedError = accountService.Withdraw(id, amount, at)
			})

			It("Should show up in account repo", func() {
				Expect(resultedError).NotTo(HaveOccurred())

				saveCalledOnce := len(accountRepo.SaveAccountCalls()) == 1
				Expect(saveCalledOnce).To(BeTrue(), `Save func of repo should be called once`)

				expectedAccount := NewAccount(id)
				expectedAccount.Withdraw(amount, at)
				Expect(accountRepo.SaveAccountCalls()[0].In1).To(Equal(expectedAccount))

				Expect(len(expectedAccount.Transactions) > 0).To(BeTrue(), "no transaction added")
			})
		})
	})
}

func sampleStatementLines() []StatementLine {
	return []StatementLine{
		{Date: support.ParseDate("10-01-2012"), Credit: 1000, Debit: 0, Balance: 1000},
		{Date: support.ParseDate("13-01-2012"), Credit: 2000, Debit: 0, Balance: 3000},
		{Date: support.ParseDate("14-01-2012"), Credit: 0, Debit: 500, Balance: 2500},
	}
}

func expectedStatementLineStrings() []string {
	return []string{
		"10/01/2012 || 1000.00 || || 1000.00",
		"13/01/2012 || 2000.00 || || 3000.00",
		"14/01/2012 || || 500.00 || 2500.00",
	}
}
