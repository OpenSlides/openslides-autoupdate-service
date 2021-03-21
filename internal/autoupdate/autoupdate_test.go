package autoupdate_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

func TestLive(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	ds := test.NewMockDatastore(closed, map[string]string{
		"collection/1/foo": `"Foo Value"`,
		"collection/1/bar": `"Bar Value"`,
	})
	s := autoupdate.New(ds, new(test.MockRestricter), test.UserUpdater{}, closed)
	kb := test.KeysBuilder{K: []string{"collection/1/foo", "collection/1/bar"}}

	buf := new(bytes.Buffer)
	w := lineWriter{maxLines: 1, wr: buf}
	if err := s.Live(context.Background(), 1, w, kb); err != nil {
		if !errors.Is(err, errWriterFull) {
			t.Fatalf("Live returned unexpected error: %v", err)
		}
	}

	expect := `{"collection/1/bar":"Bar Value","collection/1/foo":"Foo Value"}` + "\n"
	if buf.String() != expect {
		t.Errorf("Got %s, expected %s", buf.String(), expect)
	}
}

var errWriterFull = errors.New("first line full")

// lineWriter fails after the first newline
type lineWriter struct {
	maxLines int
	wr       io.Writer
	count    int
}

func (w lineWriter) Write(p []byte) (int, error) {
	if w.count >= w.maxLines {
		return 0, errWriterFull
	}

	idx := bytes.IndexByte(p, '\n')
	if idx != -1 {
		w.count++
		n, err := w.wr.Write(p[:idx+1])
		if err != nil {
			return n, err
		}
		return n, errWriterFull
	}

	return w.wr.Write(p)
}
