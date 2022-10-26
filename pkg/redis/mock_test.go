package redis_test

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/redis"
)

func getRedis() *redis.Redis {
	var c redis.Connection = mockConn{}
	if useRealRedis {
		c, _ = redis.NewConn(environment.ForTests{})
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

func (c mockConn) XREAD(ctx context.Context, count, stream, lastID string) (interface{}, error) {
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

func cmpMap(one, two map[datastore.Key][]byte) bool {
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
