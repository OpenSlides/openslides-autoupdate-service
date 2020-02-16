package topic_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate/topic"
)

func benchmarkAddWithXReceivers(count int, b *testing.B) {
	done := make(chan struct{}, 0)
	top := topic.New(topic.WithClosed(done))
	for i := 0; i < count; i++ {
		// starts a receiver that listens to the topic until an empty list is returned (done is closed)
		go func() {
			var tid uint64
			var v []string
			for {
				tid, v, _ = top.Get(context.Background(), tid)
				if len(v) == 0 {
					return
				}
			}
		}()
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		top.Add("value")
	}
	close(done)
}

func BenchmarkAddWithXReceivers1(b *testing.B)     { benchmarkAddWithXReceivers(1, b) }
func BenchmarkAddWithXReceivers10(b *testing.B)    { benchmarkAddWithXReceivers(10, b) }
func BenchmarkAddWithXReceivers100(b *testing.B)   { benchmarkAddWithXReceivers(100, b) }
func BenchmarkAddWithXReceivers1000(b *testing.B)  { benchmarkAddWithXReceivers(1_000, b) }
func BenchmarkAddWithXReceivers10000(b *testing.B) { benchmarkAddWithXReceivers(10_000, b) }

func benchmarkReadBigTopic(count int, b *testing.B) {
	top := topic.New()
	for i := 0; i < count; i++ {
		top.Add("value" + strconv.Itoa(i))
	}
	ctx := context.Background()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		top.Get(ctx, 0)
	}
}

func BenchmarkReadBigTopic1(b *testing.B)      { benchmarkReadBigTopic(1, b) }
func BenchmarkReadBigTopic10(b *testing.B)     { benchmarkReadBigTopic(10, b) }
func BenchmarkReadBigTopic100(b *testing.B)    { benchmarkReadBigTopic(100, b) }
func BenchmarkReadBigTopic1000(b *testing.B)   { benchmarkReadBigTopic(1_000, b) }
func BenchmarkReadBigTopic10000(b *testing.B)  { benchmarkReadBigTopic(10_000, b) }
func BenchmarkReadBigTopic100000(b *testing.B) { benchmarkReadBigTopic(100_000, b) }

func benchmarkReadLastBigTopic(count int, b *testing.B) {
	top := topic.New()
	for i := 0; i < count; i++ {
		top.Add("value" + strconv.Itoa(i))
	}
	tid := top.LastID()
	ctx := context.Background()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		top.Get(ctx, tid-1)
	}
}

func BenchmarkReadLastBigTopic1(b *testing.B)      { benchmarkReadLastBigTopic(1, b) }
func BenchmarkReadLastBigTopic10(b *testing.B)     { benchmarkReadLastBigTopic(10, b) }
func BenchmarkReadLastBigTopic100(b *testing.B)    { benchmarkReadLastBigTopic(100, b) }
func BenchmarkReadLastBigTopic1000(b *testing.B)   { benchmarkReadLastBigTopic(1_000, b) }
func BenchmarkReadLastBigTopic10000(b *testing.B)  { benchmarkReadLastBigTopic(10_000, b) }
func BenchmarkReadLastBigTopic100000(b *testing.B) { benchmarkReadLastBigTopic(100_000, b) }
