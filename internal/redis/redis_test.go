package redis_test

import (
	"encoding/json"
	"sort"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/redis"
	"github.com/openslides/openslides-autoupdate-service/internal/redis/conn"
)

const useRealRedis = false

func getRedis() *redis.Service {
	var c redis.Connection = mockConn{}
	if useRealRedis {
		c = conn.New("localhost:6379")
	}
	return &redis.Service{Conn: c}
}

func TestKeysChangedOnce(t *testing.T) {
	keys, err := getRedis().KeysChanged()
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	expect := []string{"key1", "key2", "key3"}
	if !cmpSlice(keys, expect) {
		t.Errorf("Expected %v, got %v", expect, keys)
	}
}

func TestKeysChangedTwice(t *testing.T) {
	r := getRedis()
	if _, err := r.KeysChanged(); err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	keys, err := r.KeysChanged()
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	expect := []string{}
	if !cmpSlice(keys, expect) {
		t.Errorf("Expected %v, got %v", expect, keys)
	}
}

type mockConn struct{}

var testData = map[string]string{
	"$": `[
		[
			"stream1",
			[
				[
					"12345-0",
					["updated", "key1", "updated", "key2"]
				],
				[
					"12346-0",
					["updated", "key1", "updated", "key3"]
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
					["updated", "key1", "updated", "key3"]
				]
			]
		]
	]`,
}

func (c mockConn) XREAD(count, block, stream, lastID string) (interface{}, error) {
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
