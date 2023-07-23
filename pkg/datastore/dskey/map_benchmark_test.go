package dskey_test

import (
	"fmt"
	"strings"
	"testing"
)

func MyBenchmark[Key comparable](b *testing.B, buildKey func(id int) Key) {
	count := 100_000
	myMap := make(map[Key]string, count)
	keys := make([]Key, count)
	for i := 0; i < count; i++ {
		key := buildKey(i)
		myMap[key] = "hello"
		keys[i] = key
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		for _, key := range keys {
			if myMap[key] != "hello" {
				b.Fatalf("value is %s, expected hello", myMap[key])
			}
		}
	}
}

func BenchmarkString(b *testing.B) {
	type Key string

	buildKey := func(id int) Key {
		return Key(fmt.Sprintf("foo/%d/bar", id))
	}

	MyBenchmark(b, buildKey)
}

func BenchmarkOldKey(b *testing.B) {
	type Key struct {
		Collection string
		ID         int
		Field      string
	}

	buildKey := func(id int) Key {
		return Key{Collection: "foo", ID: id, Field: "bar"}
	}

	MyBenchmark(b, buildKey)
}

func BenchmarkKeyWithIndex(b *testing.B) {
	type Key struct {
		value string
		idx1  int
		idx2  int
		id    int
	}

	buildKey := func(id int) Key {
		value := fmt.Sprintf("foo/%d/bar", id)
		return Key{
			value: value,
			idx1:  3,
			idx2:  strings.LastIndexByte(value, '/'),
			id:    id,
		}
	}

	MyBenchmark(b, buildKey)
}

func BenchmarkSeparateID(b *testing.B) {
	type Key struct {
		value string
		idx   int
		id    int
	}

	buildKey := func(id int) Key {
		return Key{
			value: "foo/bar",
			idx:   3,
			id:    id,
		}
	}

	MyBenchmark(b, buildKey)
}
