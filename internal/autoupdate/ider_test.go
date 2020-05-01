package autoupdate_test

import (
	"context"
	"strings"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

func TestRestrictedIDs(t *testing.T) {
	datastore := test.NewMockDatastore()
	defer datastore.Close()
	s := autoupdate.New(datastore, new(test.MockRestricter))
	defer s.Close()
	ider := s.IDer(1)

	tc := []struct {
		key     string
		idCount int
		err     string
	}{
		{"motion/1/category_ids", 2, ""},
	}
	for _, tt := range tc {
		t.Run(tt.key, func(t *testing.T) {
			ids, err := ider.IDList(context.Background(), tt.key)
			if tt.err != "" {
				if err == nil {
					t.Fatal("Got no error, expected one.")
				}
				if got := err.Error(); !strings.HasPrefix(got, tt.err) {
					t.Errorf("Got error message `%s`, expected prefix `%s`", got, tt.err)
				}
				return
			}
			if err != nil {
				t.Errorf("ider.IDList returned the unexpected error: %v", err)
			}
			if len(ids) != tt.idCount {
				t.Errorf("Got %v, expected %v", ids, tt.idCount)
			}
		})
	}
}
