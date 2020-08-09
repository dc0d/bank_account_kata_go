package account

import (
	"sort"
	"strings"

	"github.com/dc0d/bank_account_kata_go/test/support"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func DescribeStatement() {
	Describe("Statement", func() {
		DescribeTable("StatementLine to string",
			func(line StatementLine, expectedString string) {
				Expect(line.String()).To(Equal(expectedString))
			},
			Entry("1", sampleStatementLines()[0], expectedStatementLineStrings()[0]),
			Entry("2", sampleStatementLines()[1], expectedStatementLineStrings()[1]),
			Entry("3", sampleStatementLines()[2], expectedStatementLineStrings()[2]),
		)

		var (
			statement Statement
		)

		Context("Statement text output", func() {
			It("Should match the expected statement", func() {
				for _, line := range sampleStatementLines() {
					statement.AddStatementLine(line)
				}

				fromNewest := expectedStatementLineStrings()
				sort.Sort(sort.Reverse(sort.StringSlice(fromNewest)))

				lines := []string{"date || credit || debit || balance"}
				lines = append(lines, fromNewest...)
				expectedStatement := strings.Join(lines, "\n")

				Expect(statement.String()).To(Equal(expectedStatement))
			})
		})
	})
}

func sampleStatementLines() []StatementLine {
	return []StatementLine{
		{Date: support.ParseDate("10-01-2012"), Credit: 1000, Debit: 0, Balance: 1000},
		{Date: support.ParseDate("13-01-2012"), Credit: 2000, Debit: 0, Balance: 3000},
		{Date: support.ParseDate("14-01-2012"), Credit: 0, Debit: 500, Balance: 2500},
	}
}

func expectedStatementLineStrings() []string {
	return []string{
		"10/01/2012 || 1000.00 || || 1000.00",
		"13/01/2012 || 2000.00 || || 3000.00",
		"14/01/2012 || || 500.00 || 2500.00",
	}
}
