package redis_test

import (
	"bytes"
	"encoding/json"
	"sort"

	"github.com/openslides/openslides-autoupdate-service/pkg/redis"
)

func getRedis() *redis.Redis {
	var c redis.Connection = mockConn{}
	if useRealRedis {
		c = redis.NewConnection("localhost:6379")
	}
	return &redis.Redis{Conn: c}
}

type mockConn struct {
	err error
}

var testData = map[string]string{
	"$": `[
		[
			"stream1",
			[
				[
					"12345-0",
					["user/1/name", "Helga", "user/2/name", "Isolde"]
				],
				[
					"12346-0",
					["user/1/name", "Hubert", "user/3/name", "Igor"]
				]
			]
		]
	]`,
	"12345-0": `[
		[
			"stream1",
			[
				[
					"12346-0",
					["user/1/name", "Hubert", "user/3/name", "Igor"]
				]
			]
		]
	]`,
}

func (c mockConn) XREAD(count, stream, lastID string) (interface{}, error) {
	if c.err != nil {
		return nil, c.err
	}
	if _, ok := testData[lastID]; !ok {
		return nil, nil
	}
	var data interface{}
	err := json.Unmarshal([]byte(testData[lastID]), &data)
	return data, err
}

func cmpSlice(one, two []string) bool {
	if len(one) != len(two) {
		return false
	}

	sort.Strings(one)
	sort.Strings(two)
	for i := range one {
		if one[i] != two[i] {
			return false
		}
	}
	return true
}

func cmpMap(one, two map[string]json.RawMessage) bool {
	if len(one) != len(two) {
		return false
	}

	for key := range one {
		if bytes.Compare(one[key], two[key]) != 0 {
			return false
		}
	}
	return true
}
