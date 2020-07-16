//go:generate moq -pkg blackbox -out ./deposit_action_provider_spy_test.go ./../../../../delivery/webapi DepositActionProvider:DepositActionProviderSpy
//go:generate moq -pkg blackbox -out ./account_repo_spy_test.go ./../../../../core/model/account AccountRepo:AccountRepoSpy
//go:generate moq -pkg blackbox -out ./withdrawal_action_provider_spy_test.go ./../../../../delivery/webapi WithdrawalActionProvider:WithdrawalActionProviderSpy
//go:generate moq -pkg blackbox -out ./withdrawal_action_spy_test.go ./../../../../core/actions/boundaries Action:WithdrawalActionSpy
//go:generate moq -pkg blackbox -out ./print_bank_statement_action_provider_spy_test.go ./../../../../delivery/webapi PrintBankStatementActionProvider:PrintBankStatementActionProviderSpy
//go:generate moq -pkg blackbox -out ./print_bank_statement_action_spy_test.go ./../../../../core/actions/boundaries Action:PrintBankStatementActionSpy

package blackbox

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dc0d/bank_account_kata_go/core/actions"
	"github.com/dc0d/bank_account_kata_go/core/actions/boundaries"
	"github.com/dc0d/bank_account_kata_go/core/infrastructure"
	"github.com/dc0d/bank_account_kata_go/core/model/account"
	"github.com/dc0d/bank_account_kata_go/delivery/webapi"

	"github.com/dc0d/bank_account_kata_go/test/delivery/webapispec"
	"github.com/dc0d/bank_account_kata_go/test/doubles"

	"github.com/labstack/echo/v4"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test_webapi_suite_blackbox(t *testing.T) {
	webapispec.WebAPISuite(httpAPICallerFactory)

	RegisterFailHandler(Fail)
	RunSpecs(t, "WebAPI Suite")
}

type httpAPICaller struct {
	req *http.Request
	rec *httptest.ResponseRecorder

	router http.Handler

	depositActionProvider *DepositActionProviderSpy
	depositAction         *actions.Deposit

	withdrawalActionProvider *WithdrawalActionProviderSpy
	withdrawalAction         *actions.Withdrawal

	printBankStatementActionProvider *PrintBankStatementActionProviderSpy
	printBankStatementAction         *actions.PrintBankStatement

	accountRepo *AccountRepoSpy
	fakeStorage *infrastructure.FakeStorage

	stateFull bool
}

func httpAPICallerFactory(stateFull ...bool) doubles.HTTPAPICaller {
	isStateFull := false
	if len(stateFull) > 0 {
		isStateFull = stateFull[0]
	}

	res := &httpAPICaller{stateFull: isStateFull}

	if res.stateFull {
		res.router = res.initRouter()
	}

	return res
}

func (caller *httpAPICaller) Post(path string, payload interface{}) (response string, httpStatus int, err error) {
	var (
		js []byte
	)

	caller.req = nil
	caller.rec = httptest.NewRecorder()

	if caller.router == nil {
		caller.router = caller.initRouter()
	}

	js, err = json.Marshal(payload)
	if err != nil {
		return
	}

	caller.req = httptest.NewRequest(
		http.MethodPost,
		path,
		bytes.NewBuffer(js))

	caller.req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	caller.router.ServeHTTP(caller.rec, caller.req)

	return caller.rec.Body.String(), caller.rec.Code, nil
}

func (caller *httpAPICaller) Get(path string) (response string, httpStatus int, err error) {
	caller.req = nil
	caller.rec = httptest.NewRecorder()

	if caller.router == nil {
		caller.router = caller.initRouter()
	}

	caller.req = httptest.NewRequest(
		http.MethodGet,
		path,
		nil)

	caller.req.Header.Set(echo.HeaderAcceptEncoding, echo.MIMEApplicationJSON)

	caller.router.ServeHTTP(caller.rec, caller.req)

	return caller.rec.Body.String(), caller.rec.Code, nil
}

func (caller *httpAPICaller) Verify(expectedState interface{}) {
	verificationFailed := true

	if command, ok := expectedState.(webapi.DepositCommandPayload); ok {
		Expect(command).NotTo(BeZero())

		saveAccountCalledAtLeastOnce := len(caller.accountRepo.SaveAccountCalls()) > 0
		Expect(saveAccountCalledAtLeastOnce).To(BeTrue(), `save account called at least once`)

		if !caller.stateFull {
			expectedAccount := account.NewAccount(command.AccountID)
			expectedAccount.Deposit(command.Amount, command.Time)
			Expect(caller.accountRepo.SaveAccountCalls()[0].In1).To(Equal(expectedAccount))
		}

		return
	}

	if command, ok := expectedState.(webapi.WithdrawCommandPayload); ok {
		Expect(command).NotTo(BeZero())

		saveAccountCalledAtLeastOnce := len(caller.accountRepo.SaveAccountCalls()) > 0
		Expect(saveAccountCalledAtLeastOnce).To(BeTrue(), `save account called at least once`)

		if !caller.stateFull {
			expectedAccount := account.NewAccount(command.AccountID)
			expectedAccount.Withdraw(command.Amount, command.Time)
			Expect(caller.accountRepo.SaveAccountCalls()[0].In1).To(Equal(expectedAccount))
		}

		return
	}

	if _, ok := expectedState.(string); ok {
		printBankStatementActionProviderCalledOnce := len(caller.printBankStatementActionProvider.ProvidePrintBankStatementActionCalls()) == 1
		Expect(printBankStatementActionProviderCalledOnce).To(BeTrue(), `ProvidePrintBankStatementAction should be called once`)

		return
	}

	Expect(verificationFailed).To(BeFalse(), "http api caller verification failed with: ", expectedState)
}

func (caller *httpAPICaller) initRouter() http.Handler {
	caller.fakeStorage = infrastructure.NewFakeStorage()

	caller.accountRepo = &AccountRepoSpy{}
	caller.accountRepo.SaveAccountFunc = func(in1 *account.Account) error {
		return caller.fakeStorage.SaveAccount(in1)
	}
	caller.accountRepo.FindAccountFunc = func(in1 account.AccountID) (*account.Account, error) {
		return caller.fakeStorage.FindAccount(in1)
	}

	caller.depositAction = actions.NewDeposit(account.NewAccountService(caller.accountRepo))

	caller.depositActionProvider = &DepositActionProviderSpy{}
	caller.depositActionProvider.ProvideDepositActionFunc = func() boundaries.Action {
		return caller.depositAction
	}

	caller.withdrawalAction = actions.NewWithdrawal(account.NewAccountService(caller.accountRepo))

	caller.withdrawalActionProvider = &WithdrawalActionProviderSpy{}
	caller.withdrawalActionProvider.ProvideWithdrawalActionFunc = func() boundaries.Action {
		return caller.withdrawalAction
	}

	caller.printBankStatementAction = actions.NewPrintBankStatement(
		account.NewAccountService(caller.accountRepo))

	caller.printBankStatementActionProvider = &PrintBankStatementActionProviderSpy{}
	caller.printBankStatementActionProvider.ProvidePrintBankStatementActionFunc = func() boundaries.Action {
		return caller.printBankStatementAction
	}

	var router webapi.Router
	router.WithDepositActionProvider(caller.depositActionProvider)
	router.WithWithdrawalActionProvider(caller.withdrawalActionProvider)
	router.WithPrintBankStatementActionProvider(caller.printBankStatementActionProvider)
	return router.Init()
}
