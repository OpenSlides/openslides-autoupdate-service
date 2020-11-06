package dataprovider

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/OpenSlides/openslides-permission-service/internal/definitions"
)

// ExternalDataProvider is the connection to the datastore. It returns the data
// required by the permission service.
type ExternalDataProvider interface {
	// If a field does not exist, it is not returned.
	Get(ctx context.Context, fields ...definitions.Fqfield) ([]json.RawMessage, error)
}

// DataProvider is a wrapper around permission.ExternalDataProvider that
// provides some helper functions.
type DataProvider struct {
	ctx                  context.Context
	externalDataprovider ExternalDataProvider
}

// NewDataProvider returns a new DataProvider.
func NewDataProvider(ctx context.Context, externalDataprovider ExternalDataProvider) DataProvider {
	return DataProvider{ctx, externalDataprovider}
}

func (dp DataProvider) externalGet(fields ...definitions.Fqfield) ([]json.RawMessage, error) {
	return dp.externalDataprovider.Get(dp.ctx, fields...)
}

// GetString returns the value of a string field.
func (dp DataProvider) GetString(fqfield definitions.Fqfield) (string, error) {
	fields, err := dp.externalGet(fqfield)
	if err != nil {
		return "", fmt.Errorf("GetString: %w", err)
	}

	if fields[0] == nil {
		return "", fmt.Errorf("No fqfield '%s'", fqfield)
	}

	return string(fields[0]), nil
}

// GetStringWithDefault returns a string value but returns a default value, if
// the fqfield does not exist.
func (dp DataProvider) GetStringWithDefault(fqfield definitions.Fqfield, defaultValue string) (string, error) {
	fields, err := dp.externalGet(fqfield)
	if err != nil {
		return "", fmt.Errorf("GetStringWithDefault: %w", err)
	}
	value := fields[0]
	if value == nil {
		return defaultValue, nil
	}
	return string(value), nil
}

// GetStringArrayWithDefault returns a value, that conatins a list of strings.
func (dp DataProvider) GetStringArrayWithDefault(fqfield definitions.Fqfield, defaultValue []string) ([]string, error) {
	fields, err := dp.externalGet(fqfield)
	if err != nil {
		return nil, fmt.Errorf("GetStringArrayWithDefault: %w", err)
	}

	value := fields[0]
	if value == nil {
		return defaultValue, nil
	}

	var parsedValue []string
	if err := json.Unmarshal(value, &parsedValue); nil != err {
		return nil, fmt.Errorf("The fqfield '%s' is not an array of strings: %w", fqfield, err)
	}
	return parsedValue, nil
}

// GetMany returns a list of values.
func (dp DataProvider) GetMany(fqfields []definitions.Fqfield) (definitions.FqfieldData, error) {
	result, err := dp.externalGet(fqfields...)
	if err != nil {
		return nil, fmt.Errorf("GetMany: %w", err)
	}

	converted := make(map[definitions.Fqfield]definitions.Value, len(result))
	for i, v := range result {
		converted[fqfields[i]] = v
	}
	return converted, nil
}

// Exists tells, if a fqfield exist.
//
// If an error happens, it returns false.
func (dp DataProvider) Exists(fqfield definitions.Fqfield) (bool, error) {
	fields, err := dp.externalGet(fqfield)
	if err != nil {
		return false, fmt.Errorf("Exists: %w", err)
	}

	return fields[0] != nil, nil
}

// GetInt returns an int value.
func (dp DataProvider) GetInt(fqfield definitions.Fqfield) (int, error) {
	fields, err := dp.externalGet(fqfield)
	if err != nil {
		return 0, fmt.Errorf("GetInt: %w", err)
	}

	value := fields[0]
	if value == nil {
	}

	parsedValue, err := strconv.Atoi(string(value))
	if err != nil {
		return 0, fmt.Errorf("'%s' of field '%s' is not an integer: %w", string(value), fqfield, err)
	}
	return parsedValue, nil
}

// GetIntWithDefault returns a int value or the default value.
func (dp DataProvider) GetIntWithDefault(fqfield definitions.Fqfield, defaultValue int) (int, error) {
	fields, err := dp.externalGet(fqfield)
	if err != nil {
		return 0, fmt.Errorf("GetIntWithDefault: %w", err)
	}

	value := fields[0]
	if value == nil {
		return defaultValue, nil
	}

	parsedValue, err := strconv.Atoi(string(value))
	if err != nil {
		return 0, fmt.Errorf("'%s' of field '%s' is not an integer: %w", string(value), fqfield, err)
	}
	return parsedValue, nil
}

// GetIntArray returns an array of ints.
func (dp DataProvider) GetIntArray(fqfield definitions.Fqfield) ([]int, error) {
	fields, err := dp.externalGet(fqfield)
	if err != nil {
		return nil, fmt.Errorf("GetIntArray: %w", err)
	}
	value := fields[0]
	if value == nil {
		return nil, fmt.Errorf("No '%s' fqfield", fqfield)
	}

	var parsedValue []int
	if err := json.Unmarshal(value, &parsedValue); nil != err {
		return nil, fmt.Errorf("'%s' of field '%s' is not an integer array: %w", string(value), fqfield, err)
	}
	return parsedValue, nil
}

// GetIntArrayWithDefault returns an int array or the default value.
func (dp DataProvider) GetIntArrayWithDefault(fqfield definitions.Fqfield, defaultValue []int) ([]int, error) {
	fields, err := dp.externalGet(fqfield)
	if err != nil {
		return nil, fmt.Errorf("GetIntArrayWithDefault: %w", err)
	}
	value := fields[0]
	if value == nil {
		return defaultValue, nil
	}

	var parsedValue []int
	if err := json.Unmarshal(value, &parsedValue); nil != err {
		return nil, fmt.Errorf("'%s' of field '%s' is not an integer array: %w", string(value), fqfield, err)
	}
	return parsedValue, nil
}

// GetBoolWithDefault returns a bool value or the defaultValue.
func (dp DataProvider) GetBoolWithDefault(fqfield definitions.Fqfield, defaultValue bool) (bool, error) {
	fields, err := dp.externalGet(fqfield)
	if err != nil {
		return false, fmt.Errorf("GetBoolWithDefault: %w", err)
	}

	value := fields[0]
	if value == nil {
		return defaultValue, nil
	}
	var parsedValue bool
	if err := json.Unmarshal(value, &parsedValue); nil != err {
		return false, fmt.Errorf("'%s' of field '%s' is not a boolean: %w", string(value), fqfield, err)
	}
	return parsedValue, nil
}
