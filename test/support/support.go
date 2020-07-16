package support

import (
	"math/rand"
	"time"
)

func ParseDate(ddMMyyyy string) time.Time {
	t, err := time.ParseInLocation("02-01-2006", ddMMyyyy, time.Local)
	if err != nil {
		panic(err)
	}
	return t
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
