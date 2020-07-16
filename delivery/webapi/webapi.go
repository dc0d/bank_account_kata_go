package webapi

import (
	"net/http"

	"github.com/dc0d/bank_account_kata_go/core"
	"github.com/dc0d/bank_account_kata_go/core/actions/boundaries"
	"github.com/dc0d/bank_account_kata_go/core/model/account"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	DepositActionProvider
	WithdrawalActionProvider
	PrintBankStatementActionProvider

	*echo.Echo

	initialized bool
}

func (router *Router) WithDepositActionProvider(provider DepositActionProvider) {
	router.DepositActionProvider = provider
}

func (router *Router) WithWithdrawalActionProvider(provider WithdrawalActionProvider) {
	router.WithdrawalActionProvider = provider
}

func (router *Router) WithPrintBankStatementActionProvider(provider PrintBankStatementActionProvider) {
	router.PrintBankStatementActionProvider = provider
}

func (router *Router) Init() http.Handler {
	if !router.initialized {
		router.init()
		router.initialized = true
	}

	return router
}

func (router *Router) init() {
	router.Echo = echo.New()

	router.Use(middleware.Recover())

	api := router.Group("/api/bank")

	api.POST("/transactions", func(c echo.Context) error {
		commandPayload, err := extractRawCommandPayload(c)
		if err != nil {
			return err
		}

		switch commandPayload.CommandTerm() {
		case DepositCommandTerm:
			return router.handleDepositCommand(c, commandPayload)
		case WithdrawCommandTerm:
			return router.handleWithdrawCommand(c, commandPayload)
		}

		commandPayload["SERVER_ERROR"] = "COMMAND NOT HANDLED"
		return c.JSON(http.StatusBadRequest, commandPayload)
	})

	api.GET("/:account_id/statement", func(c echo.Context) error {
		var (
			accountIDParam = c.Param("account_id")

			resultedError     error
			resultedStatement account.Statement

			input = boundaries.PrintBankStatementInput(accountIDParam)

			output boundaries.PrintBankStatementOutput = func(statement account.Statement, err error) {
				resultedError = err
				resultedStatement = statement
			}

			action = router.ProvidePrintBankStatementAction()
		)

		action.Execute(input, output)

		if resultedError != nil {
			return c.String(http.StatusInternalServerError, resultedError.Error())
		}

		return c.String(http.StatusOK, resultedStatement.String())
	})
}

func (router *Router) handleDepositCommand(c echo.Context, commandPayload RawCommandPayload) error {
	var cmd DepositCommandPayload
	core.CopyData(commandPayload, &cmd)

	var (
		resultedError error

		action = router.ProvideDepositAction()

		input = boundaries.DepositInput{
			AccountID: cmd.AccountID,
			Amount:    cmd.Amount,
			Time:      cmd.Time,
		}

		output boundaries.DepositOutput = func(err error) {
			resultedError = err
		}
	)

	action.Execute(input, output)

	if resultedError != nil {
		return c.JSON(http.StatusInternalServerError, resultedError)
	}

	return c.NoContent(http.StatusOK)
}

func (router *Router) handleWithdrawCommand(c echo.Context, commandPayload RawCommandPayload) error {
	var cmd WithdrawCommandPayload
	core.CopyData(commandPayload, &cmd)

	var (
		resultedError error

		action = router.ProvideWithdrawalAction()

		input = boundaries.WithdrawalInput{
			AccountID: cmd.AccountID,
			Amount:    cmd.Amount,
			Time:      cmd.Time,
		}

		output boundaries.WithdrawalOutput = func(err error) {
			resultedError = err
		}
	)

	action.Execute(input, output)

	if resultedError != nil {
		return c.JSON(http.StatusInternalServerError, resultedError)
	}

	return c.NoContent(http.StatusOK)
}

type RawCommandPayload map[string]interface{}

func (cmd RawCommandPayload) CommandTerm() CommandTerm {
	return CommandTerm((cmd["command"]).(string))
}

func extractRawCommandPayload(c echo.Context) (RawCommandPayload, error) {
	commandPayload := make(RawCommandPayload)
	if err := c.Bind(&commandPayload); err != nil {
		return nil, err
	}
	return commandPayload, nil
}
