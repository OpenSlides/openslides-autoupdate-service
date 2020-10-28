package allowed

import (
	"fmt"

	"github.com/OpenSlides/openslides-permission-service/pkg/definitions"
)

func MakeSet(fields []definitions.Field) map[definitions.Field]bool {
	fieldMap := make(map[definitions.Field]bool)
	for _, field := range fields {
		fieldMap[field] = true
	}
	return fieldMap
}

// Returns an error, if there are fields in data, that are not in allowedFields.
func ValidateFields(data definitions.FqfieldData, allowedFields map[definitions.Field]bool) error {
	invalidFields := make([]definitions.Field, 0)
	for field, _ := range data {
		if !allowedFields[field] {
			invalidFields = append(invalidFields, field)
		}
	}

	if len(invalidFields) > 0 {
		return fmt.Errorf("Invalid fields: %v", invalidFields)
	}
	return nil
}
