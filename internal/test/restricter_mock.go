package test

import (
	"encoding/json"
	"fmt"
	"strings"
)

// MockRestricter implements the restricter interface.
type MockRestricter struct{}

// Restrict does currently nothing.
func (r *MockRestricter) Restrict(uid int, data map[string]string) {

}

func modelFormat(data map[string]json.RawMessage) (map[string]map[string]map[string]json.RawMessage, error) {
	modelData := make(map[string]map[string]map[string]json.RawMessage)
	for key, value := range data {
		keyParts := strings.SplitN(key, "/", 3)
		if len(keyParts) != 3 {
			return nil, fmt.Errorf("invalid key `%s`", key)
		}

		collection := keyParts[0]
		id := keyParts[1]
		field := keyParts[2]
		if modelData[collection] == nil {
			modelData[collection] = make(map[string]map[string]json.RawMessage)
		}
		if modelData[collection][id] == nil {
			modelData[collection][id] = make(map[string]json.RawMessage)
		}
		modelData[collection][id][field] = value
	}
	return modelData, nil
}
