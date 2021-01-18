package autoupdate_test

import (
	"context"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

func TestLive(t *testing.T) {
	datastore := new(test.MockDatastore)
	closed := make(chan struct{})
	defer close(closed)
	s := autoupdate.New(datastore, new(test.MockRestricter), test.UserUpdater{}, closed)
	kb := test.KeysBuilder{K: []string{"foo", "bar"}}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	w := new(mockWriter)
	s.Live(ctx, 1, w, kb)

	if len(w.lines) != 1 {
		t.Fatalf("Got %d lines, expected 1", len(w.lines))
	}

	expect := `{"bar":"Hello World","foo":"Hello World"}` + "\n"
	if got := w.lines[0]; got != expect {
		t.Errorf("Got %s, expected %s", got, expect)
	}
}

type mockWriter struct {
	lines []string
	buf   []byte
}

func (w *mockWriter) Write(p []byte) (int, error) {
	w.buf = append(w.buf, p...)
	return len(p), nil
}

func (w *mockWriter) Flush() {
	w.lines = append(w.lines, string(w.buf))
	w.buf = nil
}
