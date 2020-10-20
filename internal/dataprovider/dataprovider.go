package dataprovider

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/OpenSlides/openslides-permission-service/pkg/definitions"
)

type DataProvider struct {
	externalDataprovider definitions.ExternalDataProvider
}

func NewDataProvider(externalDataprovider definitions.ExternalDataProvider) DataProvider {
	return DataProvider{externalDataprovider}
}

func (dp DataProvider) GetString(fqfield definitions.Fqfield) (string, error) {
	fields := dp.externalDataprovider.Get([]string{fqfield})
	value, ok := fields[fqfield]
	if !ok {
		return "", fmt.Errorf("No " + fqfield + " key")
	}
	return value, nil
}

func (dp DataProvider) GetStringWithDefault(fqfield definitions.Fqfield, defaultValue string) string {
	fields := dp.externalDataprovider.Get([]string{fqfield})
	value, ok := fields[fqfield]
	if !ok {
		return defaultValue
	}
	return value
}

func (dp DataProvider) GetStringArrayWithDefault(fqfield definitions.Fqfield, defaultValue []string) ([]string, error) {
	fields := dp.externalDataprovider.Get([]string{fqfield})
	value, ok := fields[fqfield]
	if !ok {
		return defaultValue, nil
	}

	var parsedValue []string
	err := json.Unmarshal([]byte(value), &parsedValue)
	if nil != err {
		return nil, fmt.Errorf("The key "+fqfield+" is not an array of strings: %v", err)
	}
	return parsedValue, nil
}

func (dp DataProvider) GetMany(fqfields []definitions.Fqfield) map[definitions.Fqfield]definitions.Value {
	return dp.externalDataprovider.Get(fqfields)
}

func (dp DataProvider) Exists(fqfield definitions.Fqfield) bool {
	fields := dp.externalDataprovider.Get([]string{fqfield})

	_, ok := fields[fqfield]
	return ok
}

func (dp DataProvider) GetInt(fqfield definitions.Fqfield) (int, error) {
	fields := dp.externalDataprovider.Get([]string{fqfield})
	value, ok := fields[fqfield]
	if !ok {
		return 0, fmt.Errorf("No " + fqfield + " key")
	}

	var parsedValue int
	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf(fqfield+" is not an integer: %v", err)
	}
	return parsedValue, nil
}

func (dp DataProvider) GetIntWithDefault(fqfield definitions.Fqfield, defaultValue int) (int, error) {
	fields := dp.externalDataprovider.Get([]string{fqfield})
	value, ok := fields[fqfield]
	if !ok {
		return defaultValue, nil
	}

	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf(fqfield+" is not an integer: %v", err)
	}
	return parsedValue, nil
}

func (dp DataProvider) GetIntArray(fqfield definitions.Fqfield) ([]int, error) {
	fields := dp.externalDataprovider.Get([]string{fqfield})
	value, ok := fields[fqfield]
	if !ok {
		return nil, fmt.Errorf("No " + fqfield + " key")
	}
	var parsedValue []int
	err := json.Unmarshal([]byte(value), &parsedValue)
	if nil != err {
		return nil, fmt.Errorf(fqfield + " is not an integer array")
	}
	return parsedValue, nil
}

func (dp DataProvider) GetIntArrayWithDefault(fqfield definitions.Fqfield, defaultValue []int) ([]int, error) {
	fields := dp.externalDataprovider.Get([]string{fqfield})
	value, ok := fields[fqfield]
	if !ok {
		return defaultValue, nil
	}

	var parsedValue []int
	err := json.Unmarshal([]byte(value), &parsedValue)
	if nil != err {
		return nil, fmt.Errorf(fqfield + " is not an integer array")
	}
	return parsedValue, nil
}

func (dp DataProvider) GetBoolWithDefault(fqfield definitions.Fqfield, defaultValue bool) (bool, error) {
	fields := dp.externalDataprovider.Get([]string{fqfield})
	value, ok := fields[fqfield]
	if !ok {
		return defaultValue, nil
	}
	var parsedValue bool
	err := json.Unmarshal([]byte(value), &parsedValue)
	if nil != err {
		return false, fmt.Errorf(fqfield + " is not an boolean")
	}
	return parsedValue, nil
}
