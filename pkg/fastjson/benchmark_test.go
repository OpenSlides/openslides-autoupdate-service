package fastjson_test

import (
	"encoding/json"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/fastjson"
)

func BenchmarkFastjsonDecodeInt(b *testing.B) {
	values := [][]byte{
		[]byte("5"),
		[]byte("7"),
		[]byte("42345234"),
		[]byte("-34234"),
		[]byte("0"),
		[]byte("invalid"),
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sum := 0
		for _, v := range values {
			i, err := fastjson.DecodeInt(v)
			if err != nil {
				continue
			}
			sum += i
		}
	}
}

func BenchmarkJSONDecodeInt(b *testing.B) {
	values := [][]byte{
		[]byte("5"),
		[]byte("7"),
		[]byte("42345234"),
		[]byte("-34234"),
		[]byte("0"),
		[]byte("invalid"),
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sum := 0
		for _, v := range values {
			var i int

			if err := json.Unmarshal(v, &i); err != nil {
				continue
			}
			sum += i
		}
	}
}

func BenchmarkFastjsonDecodeIntList(b *testing.B) {
	values := [][]byte{
		[]byte("[5]"),
		[]byte("[]"),
		[]byte("[2,3,4]"),
		[]byte("[1,2,3,4,5,6,7,8,9]"),
		[]byte("invalid"),
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var sum []int
		for _, v := range values {
			list, err := fastjson.DecodeIntList(v)
			if err != nil {
				continue
			}
			sum = append(sum, list...)
		}
	}
}

func BenchmarkJSONDecodeIntList(b *testing.B) {
	values := [][]byte{
		[]byte("[5]"),
		[]byte("[]"),
		[]byte("[2,3,4]"),
		[]byte("[1,2,3,4,5,6,7,8,9]"),
		[]byte("invalid"),
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var sum []int
		for _, v := range values {
			var list []int

			if err := json.Unmarshal(v, &list); err != nil {
				continue
			}
			sum = append(sum, list...)
		}
	}
}
