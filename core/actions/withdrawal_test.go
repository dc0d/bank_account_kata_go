package actions

import (
	"errors"
	"time"

	"github.com/dc0d/bank_account_kata_go/core/actions/boundaries"
	"github.com/dc0d/bank_account_kata_go/core/model/account"

	"github.com/dc0d/bank_account_kata_go/test/support"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func DescribeWithdrawal() {
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
}
