package dsmock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"gopkg.in/yaml.v3"
)

// YAMLData creates key values from a yaml object.
//
// It is expected, that the input is a constant string. So there can not be any
// error at runtime. Therefore this function does not return an error but panics
// to get the developer a fast feetback.
func YAMLData(input string) map[dskey.Key][]byte {
	input = strings.ReplaceAll(input, "\t", "  ")

	var db map[string]interface{}
	if err := yaml.Unmarshal([]byte(input), &db); err != nil {
		panic(err)
	}

	data := make(map[dskey.Key][]byte)
	for dbKey, dbValue := range db {
		parts := strings.Split(dbKey, "/")
		switch len(parts) {
		case 1:
			map1, ok := dbValue.(map[interface{}]interface{})
			if !ok {
				panic(fmt.Errorf("invalid type in db key %s: %T", dbKey, dbValue))
			}
			for rawID, rawObject := range map1 {
				id, ok := rawID.(int)
				if !ok {
					panic(fmt.Errorf("invalid id type: got %T expected int", rawID))
				}
				field, ok := rawObject.(map[string]interface{})
				if !ok {
					panic(fmt.Errorf("invalid object type: got %T, expected map[string]interface{}", rawObject))
				}

				for fieldName, fieldValue := range field {
					key := dskey.MustKey("%s/%d/%s", dbKey, id, fieldName)
					bs, err := json.Marshal(fieldValue)
					if err != nil {
						panic(fmt.Errorf("creating test db. Key %s: %w", key, err))
					}
					data[key] = bs
				}

				idKey := dskey.MustKey("%s/%d/id", dbKey, id)
				data[idKey] = []byte(strconv.Itoa(id))
			}

		case 2:
			field, ok := dbValue.(map[string]interface{})
			if !ok {
				panic(fmt.Errorf("invalid object type: got %T, expected map[string]interface{}", dbValue))
			}

			for fieldName, fieldValue := range field {
				fqfield := dskey.MustKey("%s/%s/%s", parts[0], parts[1], fieldName)
				bs, err := json.Marshal(fieldValue)
				if err != nil {
					panic(fmt.Errorf("creating test db. Key %s: %w", fqfield, err))
				}
				data[fqfield] = bs
			}

			idKey := dskey.MustKey("%s/%s/id", parts[0], parts[1])
			data[idKey] = []byte(parts[1])

		case 3:
			key := dskey.MustKey(dbKey)
			bs, err := json.Marshal(dbValue)
			if err != nil {
				panic(fmt.Errorf("creating test db. Key %s: %w", dbKey, err))
			}

			data[key] = bs

			idKey := dskey.MustKey("%s/%s/id", parts[0], parts[1])
			data[idKey] = []byte(parts[1])
		default:
			panic(fmt.Errorf("invalid db key %s", dbKey))
		}
	}

	for k, v := range data {
		if bytes.Equal(v, []byte("null")) {
			data[k] = nil
		}
	}

	return data
}
