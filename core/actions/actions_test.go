//go:generate moq -pkg actions -out ./account_service_spy_test.go ./../model/account AccountServiceInterface:AccountServiceInterfaceSpy
//go:generate moq -pkg actions -out ./account_repo_spy_test.go ./../model/account AccountRepo:AccountRepoSpy

// a black & white test
// AccountServiceInterfaceSpy is a spy that records call inputs and outputs
// so it can be used as a mock to veryfy interaction
// this way the whitebox tests are done for designing
// also AccountServiceInterfaceSpy can call the actual AccountService and acts as a proxy
// (maybe from the testing point of view, it should be considered a decorator)
// this way the blackbox test is done

// don't do this, all this repo is just practicing and playing around with ideas
// usually whitebix and blackbox testing is about testing one thing, not cross module/package
// but I wanted to see how things will turn out

package actions

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test_actions_suite_unit(t *testing.T) {
	ActionsSuite()

	RegisterFailHandler(Fail)
	RunSpecs(t, "Actions Suite")
}

func ActionsSuite() {
	DescribeDeposit()

	DescribeWithdrawal()

	DescribePrintBankStatement()
}
