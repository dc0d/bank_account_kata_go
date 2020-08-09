package actions

import (
	"time"

	"github.com/dc0d/bank_account_kata_go/core/actions/boundaries"
	"github.com/dc0d/bank_account_kata_go/core/model/account"

	"github.com/dc0d/bank_account_kata_go/test/support"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func DescribeDeposit() {
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
}
