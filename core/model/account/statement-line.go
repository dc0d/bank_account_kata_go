package account

import (
	"strings"
	"time"
)

type StatementLine struct {
	Date    time.Time
	Credit  Amount
	Debit   Amount
	Balance Amount
}

func (sl StatementLine) String() string {
	parts := []string{
		sl.Date.Format("02/01/2006"),
		prependSpace(sl.Credit.String()),
		prependSpace(sl.Debit.String()),
		prependSpace(sl.Balance.String()),
	}
	return strings.Join(parts, " ||")
}

func prependSpace(s string) string {
	if s == "" {
		return s
	}
	return " " + s
}
