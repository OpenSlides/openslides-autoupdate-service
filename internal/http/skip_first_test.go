package http

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestSkipFirst(t *testing.T) {
	t.Run("Only one line", func(t *testing.T) {
		buf := new(bytes.Buffer)
		w := newSkipFirst(buf)

		io.Copy(w, strings.NewReader("hello world\n"))

		if got := buf.String(); got != "{}\n" {
			t.Errorf("Got `%s`, expected ``", got)
		}
	})

	t.Run("Two lines", func(t *testing.T) {
		buf := new(bytes.Buffer)
		w := newSkipFirst(buf)

		io.Copy(w, strings.NewReader("hello world\nAnd More"))

		if got := buf.String(); got != "{}\nAnd More" {
			t.Errorf("Got `%s`, expected `And More`", got)
		}
	})

	t.Run("Write first newline in part", func(t *testing.T) {
		buf := new(bytes.Buffer)
		w := newSkipFirst(buf)

		var all int
		var nCount int
		for _, step := range []string{"abc", "foo\nbar", "last"} {
			all += len(step)
			n, err := w.Write([]byte(step))
			if err != nil {
				t.Fatalf("Write: %v", err)
			}
			nCount += n
		}

		if got := buf.String(); got != "{}\nbarlast" {
			t.Errorf("Got `%s`, expected `barlast`", got)
		}

		if all != nCount {
			t.Errorf("Wrote %d bytes, expected %d", nCount, all)
		}
	})

	t.Run("Write first newline in part first", func(t *testing.T) {
		buf := new(bytes.Buffer)
		w := newSkipFirst(buf)

		var all int
		var nCount int
		for _, step := range []string{"abc", "\nfoobar", "last"} {
			all += len(step)
			n, err := w.Write([]byte(step))
			if err != nil {
				t.Fatalf("Write: %v", err)
			}
			nCount += n
		}

		if got := buf.String(); got != "{}\nfoobarlast" {
			t.Errorf("Got `%s`, expected `foobarlast`", got)
		}

		if all != nCount {
			t.Errorf("Wrote %d bytes, expected %d", nCount, all)
		}
	})

	t.Run("Write first newline in part end", func(t *testing.T) {
		buf := new(bytes.Buffer)
		w := newSkipFirst(buf)

		var all int
		var nCount int
		for _, step := range []string{"abc", "foobar\n", "last"} {
			all += len(step)
			n, err := w.Write([]byte(step))
			if err != nil {
				t.Fatalf("Write: %v", err)
			}
			nCount += n
		}

		if got := buf.String(); got != "{}\nlast" {
			t.Errorf("Got `%s`, expected `last`", got)
		}

		if all != nCount {
			t.Errorf("Wrote %d bytes, expected %d", nCount, all)
		}
	})

}
