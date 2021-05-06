package slide

import "fmt"

type slidesError struct {
	msg             string
	slide           string
	projectionID    int
	subType         string
	contentObjectID string
	meetingID       int
}

func (e slidesError) Projection() string {
	return fmt.Sprintf("projection/%d %s %s:", e.projectionID, e.contentObjectID, e.subType)
}

func (e slidesError) Error() string {
	return e.msg
}

func (e slidesError) Slide() string {
	return e.slide
}
