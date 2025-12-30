package errs

import "errors"

var (
	EndOfInput    = errors.New("end of input")
	FailedToMatch = errors.New("failed to match")
	RuleNotFound  = errors.New("rule not found")
)
