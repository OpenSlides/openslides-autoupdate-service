package slide_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/slide"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) projector.Slider {
	s := new(projector.SlideStore)
	slide.User(s)

	userSlide := s.GetSlider("user")
	assert.NotNilf(t, userSlide, "Slide with name `user` not found.")
	return userSlide
}

func TestUser(t *testing.T) {
	userSlide := setup(t)

	for _, tt := range []struct {
		name   string
		data   map[string]string
		expect string
	}{
		{
			"Only Username",
			map[string]string{
				"user/1/id":       "1",
				"user/1/username": `"jonny123"`,
			},
			`{"user":"jonny123"}`,
		},
		{
			"Only Firstname",
			map[string]string{
				"user/1/id":         "1",
				"user/1/first_name": `"Jonny"`,
			},
			`{"user":"Jonny"}`,
		},
		{
			"Only Lastname",
			map[string]string{
				"user/1/id":        "1",
				"user/1/last_name": `"Bo"`,
			},
			`{"user":"Bo"}`,
		},
		{
			"Firstname Lastname",
			map[string]string{
				"user/1/id":         "1",
				"user/1/first_name": `"Jonny"`,
				"user/1/last_name":  `"Bo"`,
			},
			`{"user":"Jonny Bo"}`,
		},
		{
			"Title Firstname Lastname",
			map[string]string{
				"user/1/id":         "1",
				"user/1/title":      `"Dr."`,
				"user/1/first_name": `"Jonny"`,
				"user/1/last_name":  `"Bo"`,
			},
			`{"user":"Dr. Jonny Bo"}`,
		},
		{
			"Title Firstname Lastname Username",
			map[string]string{
				"user/1/id":         "1",
				"user/1/username":   `"jonny123"`,
				"user/1/title":      `"Dr."`,
				"user/1/first_name": `"Jonny"`,
				"user/1/last_name":  `"Bo"`,
			},
			`{"user":"Dr. Jonny Bo"}`,
		},
		{
			"Title Username",
			map[string]string{
				"user/1/id":       "1",
				"user/1/username": `"jonny123"`,
				"user/1/title":    `"Dr."`,
			},
			`{"user":"jonny123"}`,
		},
		{
			"Title Firstname Lastname Username Level",
			map[string]string{
				"user/1/id":                       "1",
				"user/1/username":                 `"jonny123"`,
				"user/1/title":                    `"Dr."`,
				"user/1/first_name":               `"Jonny"`,
				"user/1/last_name":                `"Bo"`,
				"user/1/meeting_user_ids":         `[10,11]`,
				"meeting_user/10/structure_level": `"Bern"`,
				"meeting_user/10/meeting_id":      `222`,
				"meeting_user/11/structure_level": `"Bern-South"`,
				"meeting_user/11/meeting_id":      `223`,
			},
			`{"user":"Dr. Jonny Bo (Bern)"}`,
		},
		{
			"Title Firstname Lastname Username Level DefaultLevel",
			map[string]string{
				"user/1/id":                       "1",
				"user/1/username":                 `"jonny123"`,
				"user/1/title":                    `"Dr."`,
				"user/1/first_name":               `"Jonny"`,
				"user/1/last_name":                `"Bo"`,
				"user/1/meeting_user_ids":         `[10]`,
				"meeting_user/10/structure_level": `"Bern"`,
				"meeting_user/10/meeting_id":      `222`,
				"user/1/default_structure_level":  `"Switzerland"`,
			},
			`{"user":"Dr. Jonny Bo (Bern)"}`,
		},
		{
			"Title Firstname Lastname Username DefaultLevel",
			map[string]string{
				"user/1/id":                      "1",
				"user/1/username":                `"jonny123"`,
				"user/1/title":                   `"Dr."`,
				"user/1/first_name":              `"Jonny"`,
				"user/1/last_name":               `"Bo"`,
				"user/1/default_structure_level": `"Switzerland"`,
			},
			`{"user":"Dr. Jonny Bo (Switzerland)"}`,
		},
		{
			"Username DefaultLevel",
			map[string]string{
				"user/1/id":                      "1",
				"user/1/username":                `"jonny123"`,
				"user/1/default_structure_level": `"Switzerland"`,
			},
			`{"user":"jonny123 (Switzerland)"}`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ds := dsmock.Stub(convertData(tt.data))
			fetch := datastore.NewFetcher(ds)

			p7on := &projector.Projection{
				ContentObjectID: "user/1",
				MeetingID:       222,
			}

			bs, err := userSlide.Slide(context.Background(), fetch, p7on)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expect, string(bs))
		})
	}
}

func TestUserWithoutMeeting(t *testing.T) {
	userSlide := setup(t)

	data := convertData(map[string]string{
		"user/1/id":                      "1",
		"user/1/username":                `"jonny123"`,
		"user/1/title":                   `"Dr."`,
		"user/1/first_name":              `"Jonny"`,
		"user/1/last_name":               `"Bo"`,
		"user/1/default_structure_level": `"Switzerland"`,
	})

	ds := dsmock.Stub(data)
	fetch := datastore.NewFetcher(ds)

	p7on := &projector.Projection{
		ContentObjectID: "user/1",
	}

	bs, err := userSlide.Slide(context.Background(), fetch, p7on)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"user":"Dr. Jonny Bo (Switzerland)"}`, string(bs))
}

func TestUserWithError(t *testing.T) {
	userSlide := setup(t)
	data := map[dskey.Key][]byte{
		dskey.MustKey("user/1/id"): []byte(`1`),
	}

	ds := dsmock.Stub(data)
	fetch := datastore.NewFetcher(ds)

	p7on := &projector.Projection{
		ContentObjectID: "user/1",
		MeetingID:       222,
	}

	bs, err := userSlide.Slide(context.Background(), fetch, p7on)
	assert.Nil(t, bs)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "neither firstName, lastName nor username found")
}
