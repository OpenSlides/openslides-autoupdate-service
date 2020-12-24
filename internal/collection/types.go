package collection

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// WriteChecker is an object with the method IsAllowed.
type WriteChecker interface {
	// IsAllowed returns an error, if the given user does not have the required
	// permission for the object. If it is allowed, it can also optionaly return
	// additional data as first return parameter.
	IsAllowed(ctx context.Context, userID int, payload map[string]json.RawMessage) (map[string]interface{}, error)
}

// WriteCheckerFunc is a function with the IsAllowed signature.
type WriteCheckerFunc func(ctx context.Context, userID int, payload map[string]json.RawMessage) (map[string]interface{}, error)

// IsAllowed calls the function.
func (f WriteCheckerFunc) IsAllowed(ctx context.Context, userID int, payload map[string]json.RawMessage) (map[string]interface{}, error) {
	return f(ctx, userID, payload)
}

// ReadeChecker is an object with a method to restrict fqfields.
type ReadeChecker interface {
	RestrictFQFields(ctx context.Context, userID int, fqfields []FQField, result map[string]bool) error
}

// ReadeCheckerFunc is a function with the IsAllowed signature.
type ReadeCheckerFunc func(ctx context.Context, userID int, fqfields []FQField, result map[string]bool) error

// RestrictFQFields calls the function.
func (f ReadeCheckerFunc) RestrictFQFields(ctx context.Context, userID int, fqfields []FQField, result map[string]bool) error {
	return f(ctx, userID, fqfields, result)
}

// Connecter can connect collection.Reader and collection.Writer to the permission
// service.
type Connecter interface {
	Connect(store HandlerStore)
}

// HandlerStore can hold handlers for Readers and Writers.
type HandlerStore interface {
	RegisterReadHandler(name string, reader ReadeChecker)
	RegisterWriteHandler(name string, writer WriteChecker)
}

// FQField contains all parts of a fqfield.
type FQField struct {
	Collection string
	ID         int
	Field      string
}

// ParseFQField creates an FQField object from a fqfield string.
func ParseFQField(fqfield string) (FQField, error) {
	parts := strings.Split(fqfield, "/")
	if len(parts) != 3 {
		return FQField{}, fmt.Errorf("invalid fqfield '%s'", fqfield)
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		return FQField{}, fmt.Errorf("invalid fqfield '%s': %w", fqfield, err)
	}

	return FQField{
		Collection: parts[0],
		ID:         id,
		Field:      parts[2],
	}, nil
}

func (fqfield FQField) String() string {
	return fmt.Sprintf("%s/%d/%s", fqfield.Collection, fqfield.ID, fqfield.Field)
}
