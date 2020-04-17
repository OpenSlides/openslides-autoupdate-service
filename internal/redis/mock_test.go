package redis_test

import (
	"encoding/json"
	"sort"

	"github.com/openslides/openslides-autoupdate-service/internal/redis"
)

func getRedis() *redis.Service {
	var c redis.Connection = mockConn{}
	if useRealRedis {
		c = redis.NewConnection("localhost:6379")
	}
	return &redis.Service{Conn: c}
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
					["modified", "key1", "modified", "key2"]
				],
				[
					"12346-0",
					["modified", "key1", "modified", "key3"]
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
					["modified", "key1", "modified", "key3"]
				]
			]
		]
	]`,
}

func (c mockConn) XREAD(count, block, stream, lastID string) (interface{}, error) {
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
