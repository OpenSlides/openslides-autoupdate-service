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

	if errAdmin := ErrorForAdmin(err); errAdmin != nil {
		err = errAdmin
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
//
// This is specialiy the case for http.Do that returns the url.Error type.
//
// This function is simular to os.IsTimeout, but it uses the errors.As function
// to find sub-erros.
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

// ErrorForAdmin returns an error, if it was created with the "ForAdmin"
// function. In other case, it returns nil.
func ErrorForAdmin(err error) error {
	var errForAdmin adminError
	if errors.As(err, &errForAdmin) {
		return errForAdmin
	}
	return nil
}

type adminError struct {
	msg string
}

func (err adminError) Error() string {
	return fmt.Sprintf("ADMIN ERROR: %v", err.msg)
}

type ctxType string

const bodyCTX ctxType = "body context"

// ContextWithBody adds a body to the context.
//
// The value can be returned with the BodyFromContext function.
func ContextWithBody(ctx context.Context, body string) context.Context {
	return context.WithValue(ctx, bodyCTX, body)
}

// BodyFromContext returns the http body from a context.
func BodyFromContext(ctx context.Context) (string, bool) {
	v := ctx.Value(bodyCTX)
	if v == nil {
		return "", false
	}

	body, ok := v.(string)
	return body, ok
}

// ContextWithTag adds a tag to the context
func ContextWithTag(ctx context.Context, tag string) context.Context {
	return context.WithValue(ctx, ctxType("tag-"+tag), struct{}{})
}

// HasTagFromContext returns true if the tag was set.
func HasTagFromContext(ctx context.Context, tag string) bool {
	v := ctx.Value(ctxType("tag-" + tag))
	return v != nil
}
