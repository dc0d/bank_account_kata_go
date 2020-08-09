package account

import "strings"

type Statement struct {
	Lines []StatementLine
}

func (statement *Statement) AddStatementLine(line StatementLine) {
	statement.Lines = append(statement.Lines, line)
}

func (statement Statement) String() string {
	var sb strings.Builder
	sb.WriteString("date || credit || debit || balance\n")
	if len(statement.Lines) > 0 {
		for i := len(statement.Lines) - 1; i >= 0; i = i - 1 {
			sb.WriteString(statement.Lines[i].String() + "\n")
		}
	}
	return strings.TrimRight(sb.String(), "\n")
}
