package autoupdate_test

import (
	"context"
	"strings"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

func TestRestrictedIDs(t *testing.T) {
	keychanges := test.NewMockKeysChanged()
	defer keychanges.Close()
	s := autoupdate.New(new(test.MockRestricter), keychanges)
	defer s.Close()
	ider := s.IDer(1)

	tc := []struct {
		key     string
		idCount int
		err     string
	}{
		{"error_id", 0, "can not restrict key error_id:"},
		{"motion/1/category_ids", 2, ""},
	}
	for _, tt := range tc {
		t.Run(tt.key, func(t *testing.T) {
			ids, err := ider.IDList(context.Background(), tt.key)
			if tt.err != "" {
				if err == nil {
					t.Fatal("Expected an error, got None")
				}
				if got := err.Error(); !strings.HasPrefix(got, tt.err) {
					t.Errorf("Expected error msg to have prefix `%s`, got `%s`", tt.err, got)
				}
				return
			}
			if len(ids) != tt.idCount {
				t.Errorf("Expected %d ids, got: %v", tt.idCount, ids)
			}
		})
	}
}
