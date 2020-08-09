//go:generate moq -out ./account_repo_spy_test.go ./ AccountRepo:AccountRepoSpy

package account

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test_account_suite_unit(t *testing.T) {
	AccountSuite()

	RegisterFailHandler(Fail)
	RunSpecs(t, "Account Suite")
}

func AccountSuite() {
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

	DescribeAmount()

	DescribeStatement()

	DescribeAccountService()
}
