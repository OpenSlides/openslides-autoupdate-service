package topic_test

import (
	"context"
	"fmt"
	"sync"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate/topic"
)

func ExampleTopic() {
	top := topic.Topic{}
	var wg sync.WaitGroup

	// Start two consumers
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			ctx := context.Background()
			var tid uint64
			var values []string
			var err error
			for {
				tid, values, err = top.Get(ctx, tid)
				if err != nil {
					fmt.Printf("Did not expect an error, got: %v", err)
					return
				}
				for _, v := range values {
					fmt.Println(v)
				}
				if tid >= 2 {
					return
				}
			}
		}()
	}

	// Write the messages v1, v2 and v3
	top.Add("v1")
	top.Add("v2", "v3")
	wg.Wait()

	// Unordered output:
	// v1
	// v1
	// v2
	// v2
	// v3
	// v3
}
