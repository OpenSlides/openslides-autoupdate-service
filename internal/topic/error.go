package topic

import "strconv"

// ErrUnknownID is returned when an topic id is requested that
// is unknown in the topic.
type ErrUnknownID struct {
	First uint64
	ID    uint64
}

func (e ErrUnknownID) Error() string {
	return "id " + strconv.FormatUint(e.ID, 10) + " is unknown in topic. Lowest id is " + strconv.FormatUint(e.First, 10)
}
