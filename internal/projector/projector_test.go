package projector_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"testing"
	"time"

	"github.com/openslides/openslides-autoupdate-service/internal/projector"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
	"github.com/stretchr/testify/require"
)

func TestLiveNonExistingProjector(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Microsecond)
	defer cancel()

	p := projector.New(test.NewMockDatastore(nil), nil, closed)
	buf := new(bytes.Buffer)

	if err := p.Live(ctx, 1, firstLineWriter{wr: buf}, []int{1}); err != nil {
		if !errors.Is(err, errWriterFull) {
			t.Fatalf("Live returned unexpected error: %v", err)
		}
	}

	expect := []byte(`{"1":null}` + "\n")
	if got := buf.Bytes(); !bytes.Equal(got, expect) {
		t.Errorf("Got `%s`, expected `%s`", got, expect)
	}
}

func TestLiveExistingProjector(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Microsecond)
	defer cancel()

	ds := test.NewMockDatastore(map[string]string{
		"projector/1/current_projection_ids": "[1]",
		"projection/1/stable":                "true",
		"projection/1/content_object_id":     `"test_model/1"`,
	})
	p := projector.New(ds, nil, closed)
	buf := new(bytes.Buffer)

	if err := p.Live(ctx, 1, firstLineWriter{wr: buf}, []int{1}); err != nil {
		if !errors.Is(err, errWriterFull) {
			t.Fatalf("Live returned unexpected error: %v", err)
		}
	}

	expect := []byte(`{"1":{"1":{"data":"abc","stable":true,"content_object_id":"test_model/1"}}}` + "\n")
	require.JSONEq(t, string(expect), string(buf.Bytes()))
}

var errWriterFull = errors.New("first line full")

// firstLineWriter fails after the first newline
type firstLineWriter struct {
	wr   io.Writer
	full bool
}

func (w firstLineWriter) Write(p []byte) (int, error) {
	if w.full {
		return 0, errWriterFull
	}

	idx := bytes.IndexByte(p, '\n')
	if idx != -1 {
		w.full = true
		n, err := w.wr.Write(p[:idx+1])
		if err != nil {
			return n, err
		}
		return n, errWriterFull
	}

	return w.wr.Write(p)
}
