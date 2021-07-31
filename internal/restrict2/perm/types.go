package perm

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

// Collection is an object with a method to restrict fqfields.
type Collection interface {
	RestrictFQFields(ctx context.Context, userID int, fqfields []FQField, result map[string]bool) error
}

// CollectionFunc is a function with the Collection.RestrictFQFields signature.
type CollectionFunc func(ctx context.Context, userID int, fqfields []FQField, result map[string]bool) error

// RestrictFQFields calls the function.
func (f CollectionFunc) RestrictFQFields(ctx context.Context, userID int, fqfields []FQField, result map[string]bool) error {
	return f(ctx, userID, fqfields, result)
}

// Connecter can connect Actions and Collections to a HandlerStore.
type Connecter interface {
	Connect(store HandlerStore)
}

// ConnecterFunc is a function that implements the Connecter interface.
type ConnecterFunc func(store HandlerStore)

// Connect calls itself.
func (f ConnecterFunc) Connect(store HandlerStore) {
	f(store)
}

// HandlerStore holds collections and actions.
type HandlerStore interface {
	RegisterRestricter(name string, collection Collection)
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

// TPermission is a type of all valid permission strings.
type TPermission string
