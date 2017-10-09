package utils

import (
	"os"

	"github.com/pkg/errors"
)

// ExitError exits with err code 1
func ExitError(str string, err error) {
	ExitErrorCode(str, err, 1)
}

// ExitErrorCode exits with err code specified
func ExitErrorCode(str string, err error, code int) {
	if err != nil {
		errors.Wrap(err, str)
		os.Exit(code)
	}
}

// StringInSlice checks for string in slice
func StringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

// StringSlice makes a slice str to str pointers
func StringSlice(strs []string) []*string {
	res := make([]*string, len(strs))
	for i := 0; i < len(strs); i++ {
		res[i] = &(strs[i])
	}
	return res
}
