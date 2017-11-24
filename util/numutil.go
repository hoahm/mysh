package util

import (
	"strconv"
)

// IsDigit check whether a string is a decimal digit
func IsDigit(v string) bool {
	if _, err := strconv.Atoi(v); err == nil {
		return true
	}
	return false
}
