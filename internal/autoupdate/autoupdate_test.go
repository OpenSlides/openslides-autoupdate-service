package autoupdate_test

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/test"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/dsmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLive(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	ds := dsmock.NewMockDatastore(closed, map[string]string{
		"collection/1/foo": `"Foo Value"`,
		"collection/1/bar": `"Bar Value"`,
	})
	s := autoupdate.New(ds, test.RestrictAllowed, closed)
	kb := test.KeysBuilder{K: []string{"collection/1/foo", "collection/1/bar"}}

	w := lineWriter{maxLines: 1}
	err := s.Live(context.Background(), 1, &w, kb)

	if !errors.Is(err, errWriterFull) {
		t.Errorf("err == %q, expected errWriterFull", err)
	}

	require.Len(t, w.lines, 1)
	assert.JSONEq(t, `{"collection/1/bar":"Bar Value","collection/1/foo":"Foo Value"}`, w.lines[0])
}

func TestLiveFlushBetweenUpdates(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	ds := dsmock.NewMockDatastore(closed, map[string]string{
		"collection/1/foo": `"Foo Value"`,
		"collection/1/bar": `"Bar Value"`,
	})
	s := autoupdate.New(ds, test.RestrictAllowed, closed)
	kb := test.KeysBuilder{K: []string{"collection/1/foo", "collection/1/bar"}}

	receiving := make(chan struct{})
	w := lineWriter{maxLines: 2, received: receiving}
	done := make(chan struct{})
	var err error
	go func() {
		// Run Live in the background. It will return aerrWriterFull after two lines are written.
		err = s.Live(context.Background(), 1, &w, kb)
		close(done)
	}()

	<-receiving // Wait until the first message was received.
	ds.Send(map[string]string{"collection/1/foo": `"new data"`})
	<-receiving // Wair for the second line.
	<-done

	require.True(t, errors.Is(err, errWriterFull), "Live() returned %v, expected an errWriterFull", err)
	require.Len(t, w.lines, 2)

	assert.JSONEq(t, `{"collection/1/bar":"Bar Value","collection/1/foo":"Foo Value"}`, w.lines[0])
	assert.JSONEq(t, `{"collection/1/foo":"new data"}`, w.lines[1])
}

var errWriterFull = errors.New("first line full")

// lineWriter fails after the first newline
type lineWriter struct {
	maxLines int
	lines    []string
	received chan<- struct{}
}

func (w *lineWriter) Write(p []byte) (int, error) {
	if len(w.lines) >= w.maxLines {
		return 0, errWriterFull
	}

	idx := bytes.IndexByte(p, '\n')
	if idx != -1 {
		// Do not save the newline but add it at the first return value
		w.lines = append(w.lines, string(p[:idx]))

		if w.received != nil {
			w.received <- struct{}{}
		}

		if len(w.lines) >= w.maxLines {
			return idx, errWriterFull
		}

		return idx, nil
	}

	w.lines = append(w.lines, string(p))

	return len(p), nil
}

func (w *lineWriter) Flush() {}
