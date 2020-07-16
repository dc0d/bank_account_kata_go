//go:generate moq -pkg whitebox -out ./deposit_action_provider_spy_test.go ./../../../../delivery/webapi DepositActionProvider:DepositActionProviderSpy
//go:generate moq -pkg whitebox -out ./deposit_action_spy_test.go ./../../../../core/actions/boundaries Action:DepositActionSpy
//go:generate moq -pkg whitebox -out ./withdrawal_action_provider_spy_test.go ./../../../../delivery/webapi WithdrawalActionProvider:WithdrawalActionProviderSpy
//go:generate moq -pkg whitebox -out ./withdrawal_action_spy_test.go ./../../../../core/actions/boundaries Action:WithdrawalActionSpy
//go:generate moq -pkg whitebox -out ./print_bank_statement_action_provider_spy_test.go ./../../../../delivery/webapi PrintBankStatementActionProvider:PrintBankStatementActionProviderSpy
//go:generate moq -pkg whitebox -out ./print_bank_statement_action_spy_test.go ./../../../../core/actions/boundaries Action:PrintBankStatementActionSpy

package whitebox

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dc0d/bank_account_kata_go/core/actions/boundaries"
	"github.com/dc0d/bank_account_kata_go/core/model/account"
	"github.com/dc0d/bank_account_kata_go/delivery/webapi"

	"github.com/dc0d/bank_account_kata_go/test/delivery/webapispec"
	"github.com/dc0d/bank_account_kata_go/test/doubles"
	"github.com/dc0d/bank_account_kata_go/test/support"

	"github.com/labstack/echo/v4"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test_webapi_suite_whitebox(t *testing.T) {
	webapispec.WebAPISuite(httpAPICallerFactory)

	RegisterFailHandler(Fail)
	RunSpecs(t, "WebAPI Suite")
}

//-----------------------------------------------------------------------------

type httpAPICaller struct {
	req *http.Request
	rec *httptest.ResponseRecorder

	router http.Handler

	depositActionProvider *DepositActionProviderSpy
	depositAction         *DepositActionSpy

	withdrawalActionProvider *WithdrawalActionProviderSpy
	withdrawalAction         *WithdrawalActionSpy

	printBankStatementActionProvider *PrintBankStatementActionProviderSpy
	printBankStatementAction         *PrintBankStatementActionSpy
}

func httpAPICallerFactory(...bool) doubles.HTTPAPICaller {
	res := &httpAPICaller{}

	return res
}

func (caller *httpAPICaller) Post(path string, payload interface{}) (response string, httpStatus int, err error) {
	var (
		js []byte
	)

	caller.req = nil
	caller.rec = httptest.NewRecorder()

	caller.router = caller.initRouter()

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

	caller.router = caller.initRouter()

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
		provideDepositActionCalledOnce := len(caller.depositActionProvider.ProvideDepositActionCalls()) == 1
		Expect(provideDepositActionCalledOnce).To(BeTrue(), `ProvideDepositAction should be called once`)

		executeDepositActionOnce := len(caller.depositAction.ExecuteCalls()) == 1
		Expect(executeDepositActionOnce).To(BeTrue(), `Deposit action should be executed once`)

		expectedInput := boundaries.DepositInput{
			AccountID: command.AccountID,
			Amount:    command.Amount,
			Time:      command.Time,
		}
		Expect(caller.depositAction.ExecuteCalls()[0].Input).To(Equal(expectedInput))

		var sampleOutput boundaries.DepositOutput
		Expect(caller.depositAction.ExecuteCalls()[0].Output).To(BeAssignableToTypeOf(sampleOutput))

		return
	}

	if command, ok := expectedState.(webapi.WithdrawCommandPayload); ok {
		provideWithdrawalActionCalledOnce := len(caller.withdrawalActionProvider.ProvideWithdrawalActionCalls()) == 1
		Expect(provideWithdrawalActionCalledOnce).To(BeTrue(), `ProvideWithdrawalAction should be called once`)

		executeWithdrawActionOnce := len(caller.withdrawalAction.ExecuteCalls()) == 1
		Expect(executeWithdrawActionOnce).To(BeTrue(), `Withdrawal action should be executed once`)

		expectedInput := boundaries.WithdrawalInput{
			AccountID: command.AccountID,
			Amount:    command.Amount,
			Time:      command.Time,
		}
		Expect(caller.withdrawalAction.ExecuteCalls()[0].Input).To(Equal(expectedInput))

		var sampleOutput boundaries.WithdrawalOutput
		Expect(caller.withdrawalAction.ExecuteCalls()[0].Output).To(BeAssignableToTypeOf(sampleOutput))

		return
	}

	if _, ok := expectedState.(string); ok {
		printBankStatementActionProviderCalledOnce := len(caller.printBankStatementActionProvider.ProvidePrintBankStatementActionCalls()) == 1
		Expect(printBankStatementActionProviderCalledOnce).To(BeTrue(), `ProvidePrintBankStatementAction should be called once`)

		executePrintBankStatementActionOnce := len(caller.printBankStatementAction.ExecuteCalls()) == 1
		Expect(executePrintBankStatementActionOnce).To(BeTrue(), `Print Bank Statement action should be executed once`)

		return
	}

	Expect(verificationFailed).To(BeFalse(), "http api caller verification failed with: ", expectedState)
}

func (caller *httpAPICaller) initRouter() http.Handler {
	caller.depositAction = &DepositActionSpy{}
	caller.depositAction.ExecuteFunc = func(input interface{}, output interface{}) {}

	caller.depositActionProvider = &DepositActionProviderSpy{}
	caller.depositActionProvider.ProvideDepositActionFunc = func() boundaries.Action { return caller.depositAction }

	caller.withdrawalAction = &WithdrawalActionSpy{}
	caller.withdrawalAction.ExecuteFunc = func(input interface{}, output interface{}) {}

	caller.withdrawalActionProvider = &WithdrawalActionProviderSpy{}
	caller.withdrawalActionProvider.ProvideWithdrawalActionFunc = func() boundaries.Action { return caller.withdrawalAction }

	caller.printBankStatementAction = &PrintBankStatementActionSpy{}
	caller.printBankStatementAction.ExecuteFunc = func(input interface{}, output interface{}) {
		var (
			statement account.Statement

			lines = []account.StatementLine{
				{Date: support.ParseDate("10-01-2012"), Credit: 1000, Debit: 0, Balance: 1000},
				{Date: support.ParseDate("13-01-2012"), Credit: 2000, Debit: 0, Balance: 3000},
				{Date: support.ParseDate("14-01-2012"), Credit: 0, Debit: 500, Balance: 2500},
			}
		)

		for _, line := range lines {
			statement.AddStatementLine(line)
		}

		outputPort := output.(boundaries.PrintBankStatementOutput)
		outputPort(statement, nil)
	}

	caller.printBankStatementActionProvider = &PrintBankStatementActionProviderSpy{}
	caller.printBankStatementActionProvider.ProvidePrintBankStatementActionFunc = func() boundaries.Action { return caller.printBankStatementAction }

	var router webapi.Router
	router.WithDepositActionProvider(caller.depositActionProvider)
	router.WithWithdrawalActionProvider(caller.withdrawalActionProvider)
	router.WithPrintBankStatementActionProvider(caller.printBankStatementActionProvider)
	return router.Init()
}

//-----------------------------------------------------------------------------
