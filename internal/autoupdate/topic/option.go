package topic

// Option is an option for the topic.New() constructor.
type Option func(*Topic)

// WithClosed adds a close-channel to a topic. When the given channel is
// closed, all waiting Get()-calls get unblocked.
func WithClosed(closed <-chan struct{}) Option {
	return func(top *Topic) {
		top.closed = closed
	}
}
