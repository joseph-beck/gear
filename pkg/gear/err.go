package gear

import "errors"

var (
	ErrEndOfInput    = errors.New("gear: end of input")
	ErrFailedToMatch = errors.New("gear: failed to match")
	ErrRuleNotFound  = errors.New("gear: rule not found")
)
