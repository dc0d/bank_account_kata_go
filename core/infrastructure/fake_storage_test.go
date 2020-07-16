package infrastructure

import (
	"testing"

	"github.com/dc0d/bank_account_kata_go/core/model/account"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test_infrastructure_suite(t *testing.T) {
	InfrastructureSuite()

	RegisterFailHandler(Fail)
	RunSpecs(t, "Account Suite")
}

func InfrastructureSuite() {
	Describe("Storage", func() {
		var (
			storage *FakeStorage
		)

		BeforeEach(func() {
			storage = NewFakeStorage()
		})

		Context("When saving an account", func() {
			var (
				id            account.AccountID
				clientAccount *account.Account

				resultedError error
			)

			BeforeEach(func() {
				id = "ACCOUNT_ID"
				clientAccount = account.NewAccount(id)
			})

			BeforeEach(func() {
				resultedError = storage.SaveAccount(clientAccount)
			})

			It("Should store the account", func() {
				Expect(resultedError).NotTo(HaveOccurred())
			})
		})

		Context("When searching for an account", func() {
			var (
				id              account.AccountID
				foundAccount    *account.Account
				expectedAccount *account.Account

				resultedError error
			)

			BeforeEach(func() {
				id = "ACCOUNT_ID"

				expectedAccount = account.NewAccount(id)
				_ = storage.SaveAccount(expectedAccount)
			})

			BeforeEach(func() {
				foundAccount, resultedError = storage.FindAccount(id)
			})

			It("Should return the account", func() {
				Expect(resultedError).NotTo(HaveOccurred())
				Expect(foundAccount).To(Equal(expectedAccount))
			})
		})
	})
}

var (
	storage *FakeStorage
	_       account.AccountRepo = storage
)
