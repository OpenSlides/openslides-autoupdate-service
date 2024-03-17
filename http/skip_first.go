package http

import (
	"bytes"
	"io"
	"net/http"
)

// This is a debuggin feature. Remove me after the performance tests

// skipFirst is a wrapper around a io.Writer that replaces all the bytes before the first newline with {}
type skipFirst struct {
	io.Writer

	gotFirstNewLine bool
}

func newSkipFirst(w io.Writer) *skipFirst {
	return &skipFirst{w, false}
}

func (w *skipFirst) Write(p []byte) (int, error) {
	if w.gotFirstNewLine {
		return w.Writer.Write(p)
	}

	idx := bytes.IndexByte(p, '\n')
	if idx == -1 {
		return len(p), nil
	}

	if _, err := w.Writer.Write([]byte("{}\n")); err != nil {
		return 0, err
	}
	w.gotFirstNewLine = true

	n, err := w.Writer.Write(p[idx+1:])
	return n + idx + 1, err
}

func (w *skipFirst) Flush() {
	flusher, ok := w.Writer.(http.Flusher)
	if ok {
		flusher.Flush()
	}
}
