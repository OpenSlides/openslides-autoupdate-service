package oserror

import (
	"context"
	"errors"
	"fmt"
	"log"
)

// Handle handles an error.
//
// Ignores context closed errors.
func Handle(err error) {
	if ContextDone(err) {
		return
	}

	log.Printf("Error: %v", err)
}

// ContextDone returns true, if the given error contains a context.Canceled or
// context.DeadlineExeeded error.
func ContextDone(err error) bool {
	return errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded)
}

// Timeout tells, if the given error is a timeout error.
//
// This is true, if a wrapped error implements the method Timeout()  and it
// returns true.
func Timeout(err error) bool {
	var errTimeout interface {
		Timeout() bool
	}
	return errors.As(err, &errTimeout) && errTimeout.Timeout()
}

// ForAdmin is an error message to the admin.
func ForAdmin(format string, a ...interface{}) error {
	return adminError{
		msg: fmt.Sprintf(format, a...),
	}
}

type adminError struct {
	msg string
}

func (err adminError) Error() string {
	return fmt.Sprintf("ADMIN ERROR: %v", err.msg)
}
