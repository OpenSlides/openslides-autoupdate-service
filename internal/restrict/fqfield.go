package restrict

import (
	"fmt"
	"strconv"
	"strings"
)

// fqField contains all parts of a fqfield.
type fqField struct {
	Collection string
	ID         int
	Field      string
}

// parseFQField creates an FQField object from a fqfield string.
func parseFQField(fqfield string) (fqField, error) {
	parts := strings.Split(fqfield, "/")
	if len(parts) != 3 {
		return fqField{}, fmt.Errorf("invalid fqfield '%s'", fqfield)
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		return fqField{}, fmt.Errorf("invalid fqfield '%s': %w", fqfield, err)
	}

	return fqField{
		Collection: parts[0],
		ID:         id,
		Field:      parts[2],
	}, nil
}

func (fqfield fqField) String() string {
	return fmt.Sprintf("%s/%d/%s", fqfield.Collection, fqfield.ID, fqfield.Field)
}

// FQID returns the fqid representation of the fqfiedl.
func (fqfield fqField) FQID() string {
	return fmt.Sprintf("%s/%d", fqfield.Collection, fqfield.ID)
}

func (fqfield fqField) CollectionField() string {
	return fmt.Sprintf("%s/%s", fqfield.Collection, fqfield.Field)
}
