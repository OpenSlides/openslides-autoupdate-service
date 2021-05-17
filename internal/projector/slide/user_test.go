package slide_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/slide"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/dsmock"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) projector.Slider {
	s := new(projector.SlideStore)
	slide.User(s)

	userSlide := s.Get("user")
	assert.NotNilf(t, userSlide, "Slide with name `user` not found.")
	return userSlide
}

func TestUser(t *testing.T) {
	userSlide := setup(t)
	var addKeysExpected []string

	for _, tt := range []struct {
		name            string
		data            map[string]string
		expect          string
		addKeysExpected []string
	}{
		{
			"Only Username",
			map[string]string{
				"user/1/username": `"jonny123"`,
			},
			`{"user":"jonny123"}`,
			addKeysExpected,
		},
		{
			"Only Firstname",
			map[string]string{
				"user/1/first_name": `"Jonny"`,
			},
			`{"user":"Jonny"}`,
			addKeysExpected,
		},
		{
			"Only Lastname",
			map[string]string{
				"user/1/last_name": `"Bo"`,
			},
			`{"user":"Bo"}`,
			addKeysExpected,
		},
		{
			"Firstname Lastname",
			map[string]string{
				"user/1/first_name": `"Jonny"`,
				"user/1/last_name":  `"Bo"`,
			},
			`{"user":"Jonny Bo"}`,
			addKeysExpected,
		},
		{
			"Title Firstname Lastname",
			map[string]string{
				"user/1/title":      `"Dr."`,
				"user/1/first_name": `"Jonny"`,
				"user/1/last_name":  `"Bo"`,
			},
			`{"user":"Dr. Jonny Bo"}`,
			addKeysExpected,
		},
		{
			"Title Firstname Lastname Username",
			map[string]string{
				"user/1/username":   `"jonny123"`,
				"user/1/title":      `"Dr."`,
				"user/1/first_name": `"Jonny"`,
				"user/1/last_name":  `"Bo"`,
			},
			`{"user":"Dr. Jonny Bo"}`,
			addKeysExpected,
		},
		{
			"Title Username",
			map[string]string{
				"user/1/username": `"jonny123"`,
				"user/1/title":    `"Dr."`,
			},
			`{"user":"jonny123"}`,
			addKeysExpected,
		},
		{
			"Title Firstname Lastname Username Level",
			map[string]string{
				"user/1/username":             `"jonny123"`,
				"user/1/title":                `"Dr."`,
				"user/1/first_name":           `"Jonny"`,
				"user/1/last_name":            `"Bo"`,
				"user/1/structure_level_$":    `["222", "223"]`,
				"user/1/structure_level_$222": `"Bern"`,
				"user/1/structure_level_$223": `"Bern-South"`,
			},
			`{"user":"Dr. Jonny Bo (Bern)"}`,
			addKeysExpected,
		},
		{
			"Title Firstname Lastname Username Level DefaultLevel",
			map[string]string{
				"user/1/username":                `"jonny123"`,
				"user/1/title":                   `"Dr."`,
				"user/1/first_name":              `"Jonny"`,
				"user/1/last_name":               `"Bo"`,
				"user/1/structure_level_$":       `["222"]`,
				"user/1/structure_level_$222":    `"Bern"`,
				"user/1/default_structure_level": `"Switzerland"`,
			},
			`{"user":"Dr. Jonny Bo (Bern)"}`,
			addKeysExpected,
		},
		{
			"Title Firstname Lastname Username DefaultLevel",
			map[string]string{
				"user/1/username":                `"jonny123"`,
				"user/1/title":                   `"Dr."`,
				"user/1/first_name":              `"Jonny"`,
				"user/1/last_name":               `"Bo"`,
				"user/1/default_structure_level": `"Switzerland"`,
			},
			`{"user":"Dr. Jonny Bo (Switzerland)"}`,
			addKeysExpected,
		},
		{
			"Username DefaultLevel",
			map[string]string{
				"user/1/username":                `"jonny123"`,
				"user/1/default_structure_level": `"Switzerland"`,
			},
			`{"user":"jonny123 (Switzerland)"}`,
			addKeysExpected,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			closed := make(chan struct{})
			defer close(closed)
			ds := dsmock.NewMockDatastore(closed, tt.data)

			p7on := &projector.Projection{
				ContentObjectID: "user/1",
				MeetingID:       222,
			}

			bs, keys, err := userSlide.Slide(context.Background(), ds, p7on)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expect, string(bs))
			expectedKeys := []string{
				"user/1/username",
				"user/1/title",
				"user/1/first_name",
				"user/1/last_name",
				"user/1/structure_level_$222",
				"user/1/default_structure_level",
			}
			expectedKeys = append(expectedKeys, tt.addKeysExpected...)
			assert.ElementsMatch(t, keys, expectedKeys)
		})
	}
}

func TestUserWithoutMeeting(t *testing.T) {
	userSlide := setup(t)
	closed := make(chan struct{})
	defer close(closed)
	data := map[string]string{
		"user/1/username":                `"jonny123"`,
		"user/1/title":                   `"Dr."`,
		"user/1/first_name":              `"Jonny"`,
		"user/1/last_name":               `"Bo"`,
		"user/1/default_structure_level": `"Switzerland"`,
	}

	ds := dsmock.NewMockDatastore(closed, data)

	p7on := &projector.Projection{
		ContentObjectID: "user/1",
	}

	bs, keys, err := userSlide.Slide(context.Background(), ds, p7on)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"user":"Dr. Jonny Bo (Switzerland)"}`, string(bs))
	expectedKeys := []string{"user/1/username", "user/1/title", "user/1/first_name", "user/1/last_name", "user/1/default_structure_level"}
	assert.ElementsMatch(t, keys, expectedKeys)
}

func TestUserWithError(t *testing.T) {
	userSlide := setup(t)
	closed := make(chan struct{})
	defer close(closed)
	data := map[string]string{
		"user/1/id": `1`,
	}

	ds := dsmock.NewMockDatastore(closed, data)

	p7on := &projector.Projection{
		ContentObjectID: "user/1",
		MeetingID:       222,
	}

	bs, keys, err := userSlide.Slide(context.Background(), ds, p7on)
	assert.Nil(t, bs)
	assert.Nil(t, keys)
	assert.Error(t, err)
	fmt.Printf(err.Error())
	assert.Contains(t, err.Error(), "Neither firstName, lastName nor username found")
}
