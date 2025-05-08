package utils

import (
	"errors"
)

var ErrStrToBool = errors.New("string to bool fail")

// ParseStrBool only support "true" "1" "false" "0".
func ParseStrBool(s string) (bool, error) {
	switch s {
	case "true", "1":
		return true, nil
	case "false", "0":
		return false, nil
	default:
		return false, ErrStrToBool
	}
}

func BoolToInt64(b bool) int64 {
	if b {
		return 1
	}
	return 0
}
