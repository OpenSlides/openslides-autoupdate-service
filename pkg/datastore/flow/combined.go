package flow

import (
	"context"
	"fmt"
	"sync"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"golang.org/x/sync/errgroup"
)

type combined struct {
	defaultFlow Flow
	otherFlows  map[string]Flow
}

func (c *combined) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	grouped := make(map[string][]dskey.Key, len(keys))
	for _, key := range keys {
		if _, ok := c.otherFlows[key.CollectionField()]; ok {
			grouped[key.CollectionField()] = append(grouped[key.CollectionField()], key)
			continue
		}

		grouped["default"] = append(grouped["default"], key)
	}

	eg, ctx := errgroup.WithContext(ctx)
	dataCh := make(chan map[dskey.Key][]byte, 1)

	for flowName, keys := range grouped {
		keys := keys
		flowName := flowName
		flow, found := c.otherFlows[flowName]
		if !found {
			flow = c.defaultFlow
			flowName = "default"
		}

		eg.Go(func() error {
			values, err := flow.Get(ctx, keys...)
			if err != nil {
				return fmt.Errorf("source %s: %w", flowName, err)
			}

			dataCh <- values
			return nil
		})
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- eg.Wait()
		close(dataCh)
	}()

	result := make(map[dskey.Key][]byte, len(keys))
	for values := range dataCh {
		for k, v := range values {
			result[k] = v
		}
	}

	if err := <-errCh; err != nil {
		return nil, err
	}

	return result, nil
}

func (c *combined) Update(ctx context.Context, updateFn func(map[dskey.Key][]byte, error)) {
	cancelCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		c.defaultFlow.Update(cancelCtx, updateFn)
		cancel()
		wg.Done()
	}()

	wg.Add(len(c.otherFlows))
	for _, flow := range c.otherFlows {
		go func(flow Flow) {
			flow.Update(cancelCtx, updateFn)
			cancel()
			wg.Done()
		}(flow)
	}

	wg.Wait()
}

// Combine combines flows.
//
// One is used as default. The others are used when a corresponding key is called.
func Combine(defaultFlow Flow, keys map[string]Flow) Flow {
	return &combined{
		defaultFlow: defaultFlow,
		otherFlows:  keys,
	}
}
