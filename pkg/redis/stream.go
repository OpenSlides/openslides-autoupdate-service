package redis

import (
	"errors"
	"fmt"
)

var errNil = errors.New("nil returned")

func stream(reply interface{}, err error) (string, map[string][]byte, error) {
	// TODO Many LAST ERRORs
	if err != nil {
		return "", nil, err
	}
	if reply == nil {
		return "", nil, errNil
	}
	streams, ok := reply.([]interface{})
	if !ok {
		return "", nil, fmt.Errorf("invalid input. Data has to be a list, not %T", reply)
	}
	if len(streams) == 0 {
		return "", nil, fmt.Errorf("invalid input. No stream in data")
	}
	stream1, ok := streams[0].([]interface{})
	if !ok {
		return "", nil, fmt.Errorf("invalid input. Stream has to be a two-tuple, not %T", streams[0])
	}
	if len(stream1) != 2 {
		return "", nil, fmt.Errorf("invalid input. Stream has to be a two-tuple, got %d elements", len(stream1))
	}
	data, ok := stream1[1].([]interface{})
	if !ok {
		return "", nil, fmt.Errorf("invalid input. Stream data has to be a list, got %T", stream1[1])
	}
	var id string
	retData := make(map[string][]byte)
	for _, v := range data {
		element, ok := v.([]interface{})
		if !ok {
			return "", nil, fmt.Errorf("invalid input. Stream element has to be a two-tuple, got %T", v)
		}
		if len(element) != 2 {
			return "", nil, fmt.Errorf("invalid input. Stream element has to be a two-tuple, got %d elements", len(element))
		}
		id, ok = tostr(element[0])
		if !ok {
			return "", nil, fmt.Errorf("invalid input. Stream ID has to be a string, got %T", element[0])
		}
		kv, ok := element[1].([]interface{})
		if !ok {
			return "", nil, fmt.Errorf("invalid input. Key values has to be a list of strings, got %T", element[1])
		}
		if len(kv)%2 != 0 {
			return "", nil, fmt.Errorf("invalid input. Odd number of key value pairs")
		}
		for i := 0; i < len(kv)-1; i += 2 {
			key, ok := tostr(kv[i])
			if !ok {
				return "", nil, fmt.Errorf("invalid input. Key has to be a string, got %T", kv[i])
			}

			value, ok := toByte(kv[i+1])
			if !ok {
				return "", nil, fmt.Errorf("invalid input. Value has to be []byte, got %T", kv[i+1])
			}

			retData[key] = value
		}
	}
	return id, retData, nil
}

// autoupdateStream parses a redis autoupdateStream object to an autoupdate.KeyChanges object.
//
// The first return value is the redis autoupdateStream id. The second one is the data and
// the third is an error.
func autoupdateStream(reply interface{}, err error) (string, map[string][]byte, error) {
	id, data, err := stream(reply, err)
	if err != nil {
		return "", nil, err
	}
	return id, data, nil
}

// logoutStream parses a redis logoutStream object to an list of sessionsIDs.
//
// The first return value is the redis autoupdateStream id. The second one is the data and
// the third is an error.
func logoutStream(reply interface{}, err error) (string, []string, error) {
	id, data, err := stream(reply, err)
	if err != nil {
		return "", nil, err
	}

	var sessionIDs []string
	for key, value := range data {
		if key != "sessionId" {
			continue
		}

		sessionIDs = append(sessionIDs, string(value))
	}
	return id, sessionIDs, nil
}

// tostr converts an interface with value string or []byte to string this is an
// helper, because the test-code generates strings but the redis code generates
// []bytes.
func tostr(i interface{}) (string, bool) {
	switch rid := i.(type) {
	case string:
		return rid, true
	case []byte:
		return string(rid), true
	default:
		return "", false
	}
}

// toByte converts an interface with value string or []byte to []byte this is an
// helper, because the test-code generates strings but the redis code generates
// []bytes.
func toByte(i interface{}) ([]byte, bool) {
	switch rid := i.(type) {
	case string:
		return []byte(rid), true
	case []byte:
		return rid, true
	default:
		return nil, false
	}
}
