package webapispec

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dc0d/bank_account_kata_go/core/model/account"
	"github.com/dc0d/bank_account_kata_go/delivery/webapi"

	"github.com/dc0d/bank_account_kata_go/test/support"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func WebAPISuite(callerFactory func(stateFull ...bool) support.HTTPAPICaller) {
	Describe("Deposit Action", func() {
		var (
			caller support.HTTPAPICaller

			actualHTTPCallResponse string
			actualHTTPCallStatus   int
			actualHTTPCallErr      error
		)

		BeforeEach(func() {
			caller = callerFactory()
		})

		Context("When the client makes a deposit", func() {
			var (
				accountID account.AccountID
				amount    account.Amount
				at        time.Time

				depositCommand webapi.DepositCommandPayload
			)

			BeforeEach(func() {
				accountID = "ACCOUNT_1"
				amount = 1000
				at = support.ParseDate("10-01-2012")

				depositCommand.Init()
				depositCommand.AccountID = accountID
				depositCommand.Amount = amount
				depositCommand.Time = at
			})

			BeforeEach(func() {
				actualHTTPCallResponse, actualHTTPCallStatus, actualHTTPCallErr =
					caller.Post("/api/bank/transactions", depositCommand)
			})

			It("Should deposit the amount in bank account", func() {
				Expect(actualHTTPCallResponse).To(BeEmpty())
				Expect(actualHTTPCallErr).NotTo(HaveOccurred())
				Expect(actualHTTPCallStatus).To(Equal(http.StatusOK))

				caller.Verify(depositCommand)
			})
		})
	})

	Describe("Withdrawal Action", func() {
		var (
			caller support.HTTPAPICaller

			actualHTTPCallResponse string
			actualHTTPCallStatus   int
			actualHTTPCallErr      error
		)

		BeforeEach(func() {
			caller = callerFactory()
		})

		Context("When the client makes a withdrawal", func() {
			var (
				accountID account.AccountID
				amount    account.Amount
				at        time.Time

				withdrawCommand webapi.WithdrawCommandPayload
			)

			BeforeEach(func() {
				accountID = "ACCOUNT_1"
				amount = 1000
				at = support.ParseDate("10-01-2012")

				withdrawCommand.Init()
				withdrawCommand.AccountID = accountID
				withdrawCommand.Amount = amount
				withdrawCommand.Time = at
			})

			BeforeEach(func() {
				actualHTTPCallResponse, actualHTTPCallStatus, actualHTTPCallErr =
					caller.Post("/api/bank/transactions", withdrawCommand)
			})

			It("Should deposit the amount in bank account", func() {
				Expect(actualHTTPCallResponse).To(BeEmpty())
				Expect(actualHTTPCallErr).NotTo(HaveOccurred())
				Expect(actualHTTPCallStatus).To(Equal(http.StatusOK))

				caller.Verify(withdrawCommand)
			})
		})
	})

	Describe("Bank Statement Scenario", func() {
		var (
			expectedBankStatement string
		)

		BeforeEach(func() {
			content, err := ioutil.ReadFile("./../expected_bank_statement")
			Expect(err).NotTo(HaveOccurred())

			expectedBankStatement = string(content)
		})

		It("Given a client makes a deposit of 1000 on 10-01-2012", func() {
			var (
				caller support.HTTPAPICaller

				actualHTTPCallResponse string
				actualHTTPCallStatus   int
				actualHTTPCallErr      error

				depositCommand  webapi.DepositCommandPayload
				withdrawCommand webapi.WithdrawCommandPayload

				accountID account.AccountID
			)

			caller = callerFactory(true)

			accountID = "ACCOUNT_1"
			{
				depositCommand.Init()
				depositCommand.AccountID = accountID
				depositCommand.Amount = 1000
				depositCommand.Time = support.ParseDate("10-01-2012")

				actualHTTPCallResponse, actualHTTPCallStatus, actualHTTPCallErr =
					caller.Post("/api/bank/transactions", depositCommand)

				Expect(actualHTTPCallResponse).To(BeEmpty())
				Expect(actualHTTPCallErr).NotTo(HaveOccurred())
				Expect(actualHTTPCallStatus).To(Equal(http.StatusOK))
			}

			By("And a deposit of 2000 on 13-01-2012")
			{
				depositCommand.Init()
				depositCommand.AccountID = accountID
				depositCommand.Amount = 2000
				depositCommand.Time = support.ParseDate("13-01-2012")

				actualHTTPCallResponse, actualHTTPCallStatus, actualHTTPCallErr =
					caller.Post("/api/bank/transactions", depositCommand)

				Expect(actualHTTPCallResponse).To(BeEmpty())
				Expect(actualHTTPCallErr).NotTo(HaveOccurred())
				Expect(actualHTTPCallStatus).To(Equal(http.StatusOK))
			}

			By("And a withdrawal of 500 on 14-01-2012")
			{
				withdrawCommand.Init()
				withdrawCommand.AccountID = accountID
				withdrawCommand.Amount = 500
				withdrawCommand.Time = support.ParseDate("14-01-2012")

				actualHTTPCallResponse, actualHTTPCallStatus, actualHTTPCallErr =
					caller.Post("/api/bank/transactions", withdrawCommand)

				Expect(actualHTTPCallResponse).To(BeEmpty())
				Expect(actualHTTPCallErr).NotTo(HaveOccurred())
				Expect(actualHTTPCallStatus).To(Equal(http.StatusOK))
			}

			By("When she prints her bank statement")
			{
				path := fmt.Sprintf("/api/bank/%v/statement", accountID)
				actualHTTPCallResponse, actualHTTPCallStatus, actualHTTPCallErr =
					caller.Get(path)

				Expect(actualHTTPCallErr).NotTo(HaveOccurred())
				Expect(actualHTTPCallStatus).To(Equal(http.StatusOK), actualHTTPCallResponse)
			}

			By("Then she would see the expected bank statement")
			Expect(actualHTTPCallResponse).To(ContainSubstring(expectedBankStatement))

			caller.Verify(actualHTTPCallResponse)
		})
	})
}
