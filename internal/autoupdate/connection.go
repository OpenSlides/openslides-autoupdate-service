package autoupdate

import (
	"context"
	"fmt"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// connection holds the state of a client. It has to be created by colling
// Connect() on a autoupdate.Service instance.
type connection struct {
	autoupdate *Autoupdate
	uid        int
	kb         KeysBuilder
	tid        uint64
	filter     filter
	hotkeys    map[string]bool
}

// Next returns the next data for the user.
//
// When Next is called for the first time, it does not block. In this case, it
// is possible, that it returns an empty map.
//
// On every other call, it blocks until there is new data. In this case, the map
// is never empty.
func (c *connection) Next(ctx context.Context) (map[string][]byte, error) {
	if c.filter.empty() {
		data, err := c.data(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("creating first time data: %w", err)
		}
		return data, nil
	}

	for {
		// Blocks until the topic is closed (on server exit) or the context is done.
		tid, changedKeys, err := c.autoupdate.topic.Receive(ctx, c.tid)
		if err != nil {
			return nil, fmt.Errorf("get updated keys: %w", err)
		}

		c.tid = tid

		var links []trace.Link
		var found bool

		for _, key := range changedKeys {
			if c.hotkeys[key] {
				found = true
			}
			if before, after, ok := Cut(key, ":"); ok && before == "span" {
				spanConfig, err := decodeSpanConfig(after)
				if err != nil {
					return nil, fmt.Errorf("decoding linked span: %w", err)
				}

				links = append(links, trace.Link{
					SpanContext: trace.NewSpanContext(spanConfig),
				})
			}
		}

		if found {
			data, err := c.data(ctx, links)
			if err != nil {
				return nil, fmt.Errorf("creating later data: %w", err)
			}
			if len(data) > 0 {
				return data, nil
			}
		}
	}
}

// data returns all values from the datastore.getter.
func (c *connection) data(ctx context.Context, linkedSpans []trace.Link) (map[string][]byte, error) {
	ctx, span := otel.Tracer("autoupdate").Start(ctx, "request data update", trace.WithLinks(linkedSpans...))
	defer span.End()

	if c.tid == 0 {
		c.tid = c.autoupdate.topic.LastID()
	}

	recorder := datastore.NewRecorder(c.autoupdate.datastore)
	restricter := c.autoupdate.restricter(recorder, c.uid)

	if err := c.kb.Update(ctx, restricter); err != nil {
		return nil, fmt.Errorf("create keys for keysbuilder: %w", err)
	}

	data, err := restricter.Get(ctx, c.kb.Keys()...)
	if err != nil {
		return nil, fmt.Errorf("get restricted data: %w", err)
	}
	c.hotkeys = recorder.Keys()

	c.filter.filter(data)

	kv := make([]string, 0, len(data))
	for k, v := range data {
		kv = append(kv, fmt.Sprintf("%s: %s", k, v))
	}

	span.SetAttributes(attribute.StringSlice("data", kv))

	return data, nil
}

// Cut from go 1.18. Remove after 1.18 is releaed.
func Cut(s, sep string) (before, after string, found bool) {
	if i := strings.Index(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, "", false
}

func decodeSpanConfig(encoded string) (trace.SpanContextConfig, error) {
	parts := strings.Split(encoded, ":")
	scc := trace.SpanContextConfig{}

	traceID, err := trace.TraceIDFromHex(parts[0])
	if err != nil {
		return scc, fmt.Errorf("decoding trace id: %w", err)
	}

	spanID, err := trace.SpanIDFromHex(parts[1])
	if err != nil {
		return scc, fmt.Errorf("decoding span id: %w", err)
	}

	traceFlags := trace.TraceFlags([]byte(parts[2])[0])

	traceState, err := trace.ParseTraceState(parts[3])
	if err != nil {
		return scc, fmt.Errorf("decoding trace state: %w", err)
	}

	remote := false
	if parts[4] == "1" {
		remote = true
	}

	scc.TraceID = traceID
	scc.SpanID = spanID
	scc.TraceFlags = traceFlags
	scc.TraceState = traceState
	scc.Remote = remote
	return scc, nil
}

func encodeSpanContext(sc trace.SpanContext) string {
	remote := "0"
	if sc.IsRemote() {
		remote = "1"
	}

	return strings.Join([]string{
		sc.TraceID().String(),
		sc.SpanID().String(),
		string([]byte{byte(sc.TraceFlags())}),
		sc.TraceState().String(),
		remote,
	}, ":")
}
