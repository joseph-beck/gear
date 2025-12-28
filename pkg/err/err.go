package err

import "errors"

var (
	EndOfInput    = errors.New("end of input")
	FailedToMatch = errors.New("failed to match")
)
