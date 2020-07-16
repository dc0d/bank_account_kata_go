// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package blackbox

import (
	"github.com/dc0d/bank_account_kata_go/core/actions/boundaries"
	"github.com/dc0d/bank_account_kata_go/delivery/webapi"
	"sync"
)

var (
	lockPrintBankStatementActionProviderSpyProvidePrintBankStatementAction sync.RWMutex
)

// Ensure, that PrintBankStatementActionProviderSpy does implement webapi.PrintBankStatementActionProvider.
// If this is not the case, regenerate this file with moq.
var _ webapi.PrintBankStatementActionProvider = &PrintBankStatementActionProviderSpy{}

// PrintBankStatementActionProviderSpy is a mock implementation of webapi.PrintBankStatementActionProvider.
//
//     func TestSomethingThatUsesPrintBankStatementActionProvider(t *testing.T) {
//
//         // make and configure a mocked webapi.PrintBankStatementActionProvider
//         mockedPrintBankStatementActionProvider := &PrintBankStatementActionProviderSpy{
//             ProvidePrintBankStatementActionFunc: func() boundaries.Action {
// 	               panic("mock out the ProvidePrintBankStatementAction method")
//             },
//         }
//
//         // use mockedPrintBankStatementActionProvider in code that requires webapi.PrintBankStatementActionProvider
//         // and then make assertions.
//
//     }
type PrintBankStatementActionProviderSpy struct {
	// ProvidePrintBankStatementActionFunc mocks the ProvidePrintBankStatementAction method.
	ProvidePrintBankStatementActionFunc func() boundaries.Action

	// calls tracks calls to the methods.
	calls struct {
		// ProvidePrintBankStatementAction holds details about calls to the ProvidePrintBankStatementAction method.
		ProvidePrintBankStatementAction []struct {
		}
	}
}

// ProvidePrintBankStatementAction calls ProvidePrintBankStatementActionFunc.
func (mock *PrintBankStatementActionProviderSpy) ProvidePrintBankStatementAction() boundaries.Action {
	if mock.ProvidePrintBankStatementActionFunc == nil {
		panic("PrintBankStatementActionProviderSpy.ProvidePrintBankStatementActionFunc: method is nil but PrintBankStatementActionProvider.ProvidePrintBankStatementAction was just called")
	}
	callInfo := struct {
	}{}
	lockPrintBankStatementActionProviderSpyProvidePrintBankStatementAction.Lock()
	mock.calls.ProvidePrintBankStatementAction = append(mock.calls.ProvidePrintBankStatementAction, callInfo)
	lockPrintBankStatementActionProviderSpyProvidePrintBankStatementAction.Unlock()
	return mock.ProvidePrintBankStatementActionFunc()
}

// ProvidePrintBankStatementActionCalls gets all the calls that were made to ProvidePrintBankStatementAction.
// Check the length with:
//     len(mockedPrintBankStatementActionProvider.ProvidePrintBankStatementActionCalls())
func (mock *PrintBankStatementActionProviderSpy) ProvidePrintBankStatementActionCalls() []struct {
} {
	var calls []struct {
	}
	lockPrintBankStatementActionProviderSpyProvidePrintBankStatementAction.RLock()
	calls = mock.calls.ProvidePrintBankStatementAction
	lockPrintBankStatementActionProviderSpyProvidePrintBankStatementAction.RUnlock()
	return calls
}
