// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package actions

import (
	"github.com/dc0d/bank_account_kata_go/core/model/account"
	"sync"
	"time"
)

var (
	lockAccountServiceInterfaceSpyDeposit            sync.RWMutex
	lockAccountServiceInterfaceSpyPrintBankStatement sync.RWMutex
	lockAccountServiceInterfaceSpyWithdraw           sync.RWMutex
)

// Ensure, that AccountServiceInterfaceSpy does implement account.AccountServiceInterface.
// If this is not the case, regenerate this file with moq.
var _ account.AccountServiceInterface = &AccountServiceInterfaceSpy{}

// AccountServiceInterfaceSpy is a mock implementation of account.AccountServiceInterface.
//
//     func TestSomethingThatUsesAccountServiceInterface(t *testing.T) {
//
//         // make and configure a mocked account.AccountServiceInterface
//         mockedAccountServiceInterface := &AccountServiceInterfaceSpy{
//             DepositFunc: func(in1 account.AccountID, in2 account.Amount, in3 time.Time) error {
// 	               panic("mock out the Deposit method")
//             },
//             PrintBankStatementFunc: func(in1 account.AccountID) (account.Statement, error) {
// 	               panic("mock out the PrintBankStatement method")
//             },
//             WithdrawFunc: func(in1 account.AccountID, in2 account.Amount, in3 time.Time) error {
// 	               panic("mock out the Withdraw method")
//             },
//         }
//
//         // use mockedAccountServiceInterface in code that requires account.AccountServiceInterface
//         // and then make assertions.
//
//     }
type AccountServiceInterfaceSpy struct {
	// DepositFunc mocks the Deposit method.
	DepositFunc func(in1 account.AccountID, in2 account.Amount, in3 time.Time) error

	// PrintBankStatementFunc mocks the PrintBankStatement method.
	PrintBankStatementFunc func(in1 account.AccountID) (account.Statement, error)

	// WithdrawFunc mocks the Withdraw method.
	WithdrawFunc func(in1 account.AccountID, in2 account.Amount, in3 time.Time) error

	// calls tracks calls to the methods.
	calls struct {
		// Deposit holds details about calls to the Deposit method.
		Deposit []struct {
			// In1 is the in1 argument value.
			In1 account.AccountID
			// In2 is the in2 argument value.
			In2 account.Amount
			// In3 is the in3 argument value.
			In3 time.Time
		}
		// PrintBankStatement holds details about calls to the PrintBankStatement method.
		PrintBankStatement []struct {
			// In1 is the in1 argument value.
			In1 account.AccountID
		}
		// Withdraw holds details about calls to the Withdraw method.
		Withdraw []struct {
			// In1 is the in1 argument value.
			In1 account.AccountID
			// In2 is the in2 argument value.
			In2 account.Amount
			// In3 is the in3 argument value.
			In3 time.Time
		}
	}
}

// Deposit calls DepositFunc.
func (mock *AccountServiceInterfaceSpy) Deposit(in1 account.AccountID, in2 account.Amount, in3 time.Time) error {
	if mock.DepositFunc == nil {
		panic("AccountServiceInterfaceSpy.DepositFunc: method is nil but AccountServiceInterface.Deposit was just called")
	}
	callInfo := struct {
		In1 account.AccountID
		In2 account.Amount
		In3 time.Time
	}{
		In1: in1,
		In2: in2,
		In3: in3,
	}
	lockAccountServiceInterfaceSpyDeposit.Lock()
	mock.calls.Deposit = append(mock.calls.Deposit, callInfo)
	lockAccountServiceInterfaceSpyDeposit.Unlock()
	return mock.DepositFunc(in1, in2, in3)
}

// DepositCalls gets all the calls that were made to Deposit.
// Check the length with:
//     len(mockedAccountServiceInterface.DepositCalls())
func (mock *AccountServiceInterfaceSpy) DepositCalls() []struct {
	In1 account.AccountID
	In2 account.Amount
	In3 time.Time
} {
	var calls []struct {
		In1 account.AccountID
		In2 account.Amount
		In3 time.Time
	}
	lockAccountServiceInterfaceSpyDeposit.RLock()
	calls = mock.calls.Deposit
	lockAccountServiceInterfaceSpyDeposit.RUnlock()
	return calls
}

// PrintBankStatement calls PrintBankStatementFunc.
func (mock *AccountServiceInterfaceSpy) PrintBankStatement(in1 account.AccountID) (account.Statement, error) {
	if mock.PrintBankStatementFunc == nil {
		panic("AccountServiceInterfaceSpy.PrintBankStatementFunc: method is nil but AccountServiceInterface.PrintBankStatement was just called")
	}
	callInfo := struct {
		In1 account.AccountID
	}{
		In1: in1,
	}
	lockAccountServiceInterfaceSpyPrintBankStatement.Lock()
	mock.calls.PrintBankStatement = append(mock.calls.PrintBankStatement, callInfo)
	lockAccountServiceInterfaceSpyPrintBankStatement.Unlock()
	return mock.PrintBankStatementFunc(in1)
}

// PrintBankStatementCalls gets all the calls that were made to PrintBankStatement.
// Check the length with:
//     len(mockedAccountServiceInterface.PrintBankStatementCalls())
func (mock *AccountServiceInterfaceSpy) PrintBankStatementCalls() []struct {
	In1 account.AccountID
} {
	var calls []struct {
		In1 account.AccountID
	}
	lockAccountServiceInterfaceSpyPrintBankStatement.RLock()
	calls = mock.calls.PrintBankStatement
	lockAccountServiceInterfaceSpyPrintBankStatement.RUnlock()
	return calls
}

// Withdraw calls WithdrawFunc.
func (mock *AccountServiceInterfaceSpy) Withdraw(in1 account.AccountID, in2 account.Amount, in3 time.Time) error {
	if mock.WithdrawFunc == nil {
		panic("AccountServiceInterfaceSpy.WithdrawFunc: method is nil but AccountServiceInterface.Withdraw was just called")
	}
	callInfo := struct {
		In1 account.AccountID
		In2 account.Amount
		In3 time.Time
	}{
		In1: in1,
		In2: in2,
		In3: in3,
	}
	lockAccountServiceInterfaceSpyWithdraw.Lock()
	mock.calls.Withdraw = append(mock.calls.Withdraw, callInfo)
	lockAccountServiceInterfaceSpyWithdraw.Unlock()
	return mock.WithdrawFunc(in1, in2, in3)
}

// WithdrawCalls gets all the calls that were made to Withdraw.
// Check the length with:
//     len(mockedAccountServiceInterface.WithdrawCalls())
func (mock *AccountServiceInterfaceSpy) WithdrawCalls() []struct {
	In1 account.AccountID
	In2 account.Amount
	In3 time.Time
} {
	var calls []struct {
		In1 account.AccountID
		In2 account.Amount
		In3 time.Time
	}
	lockAccountServiceInterfaceSpyWithdraw.RLock()
	calls = mock.calls.Withdraw
	lockAccountServiceInterfaceSpyWithdraw.RUnlock()
	return calls
}