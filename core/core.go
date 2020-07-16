package core

import (
	"encoding/json"
	"errors"
)

func CopyData(src interface{}, dest interface{}) {
	FromJSON(ToJSON(src), dest)
}

func ToJSON(v interface{}) []byte {
	js, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return js
}

func FromJSON(js []byte, v interface{}) {
	err := json.Unmarshal(js, v)
	if err != nil {
		panic(err)
	}
}

//

var (
	ErrNotFound = errors.New("ERR: NOT FOUND")
)
