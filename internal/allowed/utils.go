package allowed

import (
	"fmt"

	"github.com/OpenSlides/openslides-permission-service/internal/definitions"
)

// MakeSet does ...
func MakeSet(fields []definitions.Field) map[definitions.Field]bool {
	fieldMap := make(map[definitions.Field]bool)
	for _, field := range fields {
		fieldMap[field] = true
	}
	return fieldMap
}

// ValidateFields returns an error, if there are fields in data, that are not in
// allowedFields.
func ValidateFields(data definitions.FqfieldData, allowedFields map[definitions.Field]bool) error {
	invalidFields := make([]definitions.Field, 0)
	for field := range data {
		if !allowedFields[field] {
			invalidFields = append(invalidFields, field)
		}
	}

	if len(invalidFields) > 0 {
		return fmt.Errorf("Invalid fields: %v", invalidFields)
	}
	return nil
}
