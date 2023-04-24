package http

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/metric"
)

type userConnectionCounter interface {
	ConnectionAdd(ctx context.Context, uid int) error
	ConnectionDone(ctx context.Context, uid int) error
	ConnectionShow(ctx context.Context) (map[string]int, error)
}

type combinedCounter struct {
	metricCounter *metric.CurrentCounter
	redisCounter  userConnectionCounter
}

func (c combinedCounter) Add(ctx context.Context, uid int) error {
	c.metricCounter.Add()
	if err := c.redisCounter.ConnectionAdd(ctx, uid); err != nil {
		return fmt.Errorf("user connection counter: %w", err)
	}

	return nil
}

func (c combinedCounter) Done(ctx context.Context, uid int) error {
	c.metricCounter.Done()
	if err := c.redisCounter.ConnectionDone(ctx, uid); err != nil {
		return fmt.Errorf("user connection counter: %w", err)
	}

	return nil
}
