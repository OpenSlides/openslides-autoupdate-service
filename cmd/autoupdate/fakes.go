package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

// fakeReceiver implements the Receiver interface. It reads on a Reader, for example stdin and
// takes each word on each line as changed key.
type faker struct {
	test.MockRestricter
	buf *bufio.Reader
}

func (r *faker) KeysChanged() ([]string, error) {
	msg, err := r.buf.ReadString('\n')
	if err == io.EOF {
		// Don't return anything (block forever) if the reader is empty.
		select {}
	}
	if err != nil {
		return nil, fmt.Errorf("can not read from buffer: %w", err)
	}

	data := strings.Split(strings.TrimSpace(msg), " ")
	keys := []string{}
	for _, d := range data {
		keyValue := strings.SplitN(d, "=", 2)
		if len(keyValue) == 1 {
			keyValue = append(keyValue, fmt.Sprintf(`"The time is: %s"`, time.Now()))
		}
		keys = append(keys, keyValue[0])
		r.Update(map[string]string{keyValue[0]: keyValue[1]})
	}
	return keys, nil
}

// fake Auth implements the Authenticater interface. It always returns the given number.
type fakeAuth int

func (a fakeAuth) Authenticate(ctx context.Context, _ *http.Request) (int, error) {
	return int(a), nil
}
