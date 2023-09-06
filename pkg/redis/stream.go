package redis

import (
	"errors"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/gomodule/redigo/redis"
)

// parseStream parses one stream from a xread request.
//
// The provided function is valled for any field value pair in the stream.
//
// Returns the last id.
func parseStream(reply any, f func(k, v []byte)) (string, error) {
	valueList, err := redis.Values(reply, nil)
	if err != nil {
		return "", err
	}

	var lastID string
	for i, value := range valueList {
		idFields, ok := value.([]any)
		if !ok || len(idFields) != 2 {
			return "", fmt.Errorf("invalid stream value %d, got %v", i, value)
		}

		id, err := redis.String(idFields[0], nil)
		if err != nil {
			return "", fmt.Errorf("parsing id from entry %d: %w", i, err)
		}

		lastID = id

		fieldList, ok := idFields[1].([]any)
		if !ok || len(fieldList)%2 != 0 {
			return "", fmt.Errorf("invalid field list value %d, got %v", i, idFields[i])
		}

		for fi := 0; fi < len(fieldList); fi += 2 {
			key, ok := toByte(fieldList[fi])
			if !ok {
				return "", fmt.Errorf("field %d in entry %d is not a bulk string value, got %T", fi, i, fieldList[fi])
			}

			value, ok := toByte(fieldList[fi+1])
			if !ok {
				return "", fmt.Errorf("value %d in entry %d is not a bulk string value, got %T", fi+1, i, fieldList[fi])
			}

			f(key, value)
		}
	}
	return lastID, nil
}

// only Stream filters a xread request for one stream.
func onlyStream(reply any, only string, f func(k, v []byte)) (string, error) {
	streams, err := redis.Values(reply, nil)
	if err != nil {
		return "", fmt.Errorf("parsing reply: %w", err)
	}

	for i, stream := range streams {
		nameEntries, ok := stream.([]any)
		if !ok || len(nameEntries) != 2 {
			return "", errors.New("stream entry expects two value result")
		}

		name, err := redis.String(nameEntries[0], nil)
		if err != nil {
			return "", fmt.Errorf("parsing name of stream %d: %w", i, err)
		}

		if name != only {
			continue
		}

		lastID, err := parseStream(nameEntries[1], f)
		if err != nil {
			return "", fmt.Errorf("parsing entries of stream %d: %w", i, err)
		}

		return lastID, nil
	}

	return "", fmt.Errorf("stream not found")
}

func parseMessageBus(reply any) (string, map[dskey.Key][]byte, error) {
	data := make(map[dskey.Key][]byte)
	databuilder := func(k, v []byte) {
		key, err := dskey.FromString(string(k))
		if err != nil {
			// Ignore invalid keys
			return
		}

		if string(v) == "null" {
			v = nil
		}

		data[key] = v
	}

	lastID, err := onlyStream(reply, fieldChangedTopic, databuilder)
	if err != nil {
		return "", nil, fmt.Errorf("parsing autoupdate stream: %w", err)
	}

	return lastID, data, nil
}

// logoutStream parses a redis logoutStream object to an list of sessionsIDs.
//
// The first return value is the redis autoupdateStream id. The second one is the data and
// the third is an error.
func logoutStream(reply any) (string, []string, error) {
	var sessionIDs []string
	databuilder := func(k, v []byte) {
		if string(k) != "sessionId" {
			return
		}

		sessionIDs = append(sessionIDs, string(v))
	}

	lastID, err := onlyStream(reply, logoutTopic, databuilder)
	if err != nil {
		return "", nil, fmt.Errorf("parsing logout stream: %w", err)
	}

	return lastID, sessionIDs, nil
}

// toByte converts an interface with value string or []byte to []byte this is an
// helper, because the test-code generates strings but the redis code generates
// []bytes.
func toByte(i any) ([]byte, bool) {
	switch rid := i.(type) {
	case string:
		return []byte(rid), true
	case []byte:
		return rid, true
	default:
		return nil, false
	}
}
