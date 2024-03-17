package autoupdate

import "context"

type token struct{}

type workPool struct {
	sem chan token
}

func newWorkPool(limit int) *workPool {
	return &workPool{
		sem: make(chan token, limit),
	}
}

func (w *workPool) Do(ctx context.Context, f func() error) error {
	done, err := w.Wait(ctx)
	if err != nil {
		return err
	}
	defer done()

	return f()
}

func (w *workPool) Wait(ctx context.Context) (func(), error) {
	select {
	case w.sem <- token{}:
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	return func() {
		<-w.sem
	}, nil
}
