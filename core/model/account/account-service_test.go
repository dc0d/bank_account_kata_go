package account

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func DescribeAccountService() {
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
