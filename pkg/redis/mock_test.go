package redis_test

import (
	"bytes"
	"encoding/json"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/redis"
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

func (c mockConn) XADD(stream, key string, value []byte) error {
	return nil
}

func (c mockConn) ZINCR(key string, value []byte) error {
	return nil
}

func (c mockConn) ZRANGE(key string) (interface{}, error) {
	return nil, nil
}

func cmpMap(one, two map[string][]byte) bool {
	if len(one) != len(two) {
		return false
	}

	for key := range one {
		if !bytes.Equal(one[key], two[key]) {
			return false
		}
	}
	return true
}
