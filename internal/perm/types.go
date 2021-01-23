package perm

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// ActionChecker is an object with the method IsAllowed.
type ActionChecker interface {
	// IsAllowed tells, if the user has the permission for the object this
	// method is called on.
	IsAllowed(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error)
}

// ActionCheckerFunc is a function with the IsAllowed signature.
type ActionCheckerFunc func(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error)

// IsAllowed calls the function.
func (f ActionCheckerFunc) IsAllowed(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
	return f(ctx, userID, payload)
}

// RestricterChecker is an object with a method to restrict fqfields.
type RestricterChecker interface {
	RestrictFQFields(ctx context.Context, userID int, fqfields []FQField, result map[string]bool) error
}

// RestricterCheckerFunc is a function with the IsAllowed signature.
type RestricterCheckerFunc func(ctx context.Context, userID int, fqfields []FQField, result map[string]bool) error

// RestrictFQFields calls the function.
func (f RestricterCheckerFunc) RestrictFQFields(ctx context.Context, userID int, fqfields []FQField, result map[string]bool) error {
	return f(ctx, userID, fqfields, result)
}

// Connecter can connect collection.Reader and collection.Writer to the permission
// service.
type Connecter interface {
	Connect(store HandlerStore)
}

// ConnecterFunc is a function that implements the Connecter interface.
type ConnecterFunc func(store HandlerStore)

// Connect calls itself.
func (f ConnecterFunc) Connect(store HandlerStore) {
	f(store)
}

// HandlerStore can hold handlers for Readers and Writers.
type HandlerStore interface {
	RegisterRestricter(name string, reader RestricterChecker)
	RegisterAction(name string, writer ActionChecker)
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

// FQID returns the fqid representation of the fqfiedl.
func (fqfield FQField) FQID() string {
	return fmt.Sprintf("%s/%d", fqfield.Collection, fqfield.ID)
}
