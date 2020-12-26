package perm

import (
	"errors"
	"fmt"
)

// NotAllowedf is an error that sends a message to the client that indicates,
// that the user has not the required permissions.
func NotAllowedf(format string, a ...interface{}) error {
	return NotAllowedError{fmt.Sprintf(format, a...)}
}

// NotAllowedError tells, that the user does not have the required permission.
type NotAllowedError struct {
	reason string
}

func (e NotAllowedError) Error() string {
	return e.reason
}

// IsAllowed is a helper around functions, that return NotAllowedErrors
//
// If err == nil, it returned true. For errors, it filters out NotAllowedErrors.
// If err wrapps an NotAllowedError, it teturns false. For other errors, the
// error is returned.
//
// Example: perm.IsAllowed(perm.EnsurePerm(ctx, dp, userID, meetingID, "my.perm"))
func IsAllowed(err error) (bool, error) {
	if err == nil {
		return true, nil
	}

	var errNotAllowed NotAllowedError
	if errors.Is(err, &errNotAllowed) {
		return false, nil
	}
	return false, err
}
