package account

import "fmt"

type Amount int

func (amount Amount) String() string {
	if amount == 0 {
		return ""
	}
	return fmt.Sprintf("%.2f", float64(amount))
}
