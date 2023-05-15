package http

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/metric"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/oserror"
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

func (c combinedCounter) Metric(con metric.Container) {
	ctx := context.Background()

	data, err := c.redisCounter.ConnectionShow(ctx)
	if err != nil {
		oserror.Handle(fmt.Errorf("connection count metric: %w", err))
		return
	}

	currentConnections := 0
	averageCount := 0
	averageSum := 0
	for k, v := range data {
		if v <= 0 {
			continue
		}

		currentConnections++

		if k != "0" {
			averageCount++
			averageSum += v
		}
	}

	average := 0
	if averageCount > 0 {
		average = averageSum / averageCount
	}

	c.metricCounter.Metric(con)
	prefix := "overall_connected_users_"
	con.Add(prefix+"current", currentConnections)
	con.Add(prefix+"total", len(data))
	con.Add(prefix+"average", average)
	con.Add(prefix+"anonymous", data["0"])
}
