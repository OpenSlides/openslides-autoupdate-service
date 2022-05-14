package auerror

import (
	"context"
	"errors"
	"log"
)

// Handle handles an error.
//
// Ignores context closed errors.
func Handle(err error) {
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return
	}

	log.Printf("Error: %v", err)
}
