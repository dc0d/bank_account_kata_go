package account

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func DescribeAmount() {
	Describe("Amount", func() {
		DescribeTable("check all room types",
			func(amount Amount, expectedString string) {
				actualString := fmt.Sprint(amount)

				Expect(actualString).To(Equal(expectedString))
			},
			Entry("500", Amount(500), "500.00"),
			Entry("0", Amount(0), ""),
		)
	})
}
