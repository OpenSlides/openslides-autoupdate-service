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

// faker implements the Datastore interface. It reads form a Reader, for example
// stdin and takes each word on each line as changed key.
//
// If it is created with newFaker(), it starts an fake datastore server. The nil
// value can also be used but does nothing.
type faker struct {
	ts  *test.DatastoreServer
	buf *bufio.Reader
}

func newFaker(r io.Reader) *faker {
	f := new(faker)

	// This starts the fake datastore service.
	f.ts = test.NewDatastoreServer()

	f.buf = bufio.NewReader(r)
	return f
}

// KeysChanged blocks, until there in new data. The nil value blocks forever. If
// the faker was initialized with a reader, it reads each line form it and
// interpretes each word (separated by space) and a key that should be updated.
func (f *faker) KeysChanged() ([]string, error) {
	if f == nil {
		// If the faker was not initualized. Block forever.
		select {}
	}

	msg, err := f.buf.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			// Don't return anything (block forever) if the reader is empty.
			select {}
		}
		return nil, fmt.Errorf("read from buffer: %w", err)
	}

	data := strings.Split(strings.TrimSpace(msg), " ")
	var keys []string
	for _, d := range data {
		keyValue := strings.SplitN(d, "=", 2)
		if len(keyValue) == 1 {
			keyValue = append(keyValue, fmt.Sprintf(`"The time is: %s"`, time.Now()))
		}
		keys = append(keys, keyValue[0])
		f.ts.Update(map[string]string{keyValue[0]: keyValue[1]})
	}
	return keys, nil
}

// fake Auth implements the Authenticater interface. It always returns the given number.
type fakeAuth int

func (a fakeAuth) Authenticate(context.Context, *http.Request) (int, error) {
	return int(a), nil
}
