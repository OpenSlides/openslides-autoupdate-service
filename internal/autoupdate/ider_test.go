package autoupdate_test

import (
	"context"
	"strings"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
)

func TestRestrictedIDs(t *testing.T) {
	keychanges := newMockKeyChanged()
	defer keychanges.close()
	s := autoupdate.New(MockRestricter{}, keychanges)
	ider := s.IDer(1)

	tc := []struct {
		name    string
		key     string
		idCount int
		err     string
	}{
		{"No Reference", "motion/1/name", 0, "key motion/1/name can not be a reference; expected suffix _id or _ids"},
		{"Restricter error", "error_id", 0, "can not restrict key error_id:"},
		{"ID field", "motion/1/category_id", 1, ""},
		{"IDs field", "motion/1/category_ids", 2, ""},
	}
	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			ids, err := ider.IDs(context.Background(), tt.key)
			if tt.err != "" {
				if err == nil {
					t.Fatal("Expected an error, got None")
				}
				if got := err.Error(); !strings.HasPrefix(got, tt.err) {
					t.Errorf("Expected error msg to be `%s`, got `%s`", tt.err, got)
				}
				return
			}
			if len(ids) != tt.idCount {
				t.Errorf("Expected %d ids, got: %v", tt.idCount, ids)
			}
		})
	}
}

func TestRestrictedIDsListErrors(t *testing.T) {
	tc := []struct {
		name  string
		value string
		err   string
	}{
		{"Empty", "", "invalid value, expect list of ints"},
		{"Single number", "123", "expected first and last byte to be [ and ]"},
		{"Letters", "abc", "expected first and last byte to be [ and ]"},
		{"Not closing", "[1,2,3", "expected first and last byte to be [ and ]"},
		{"Not starting", "1,2,3]", "expected first and last byte to be [ and ]"},
		{"With Whilespace", "[1, 2, 3]", "can not convert value ` 2` to int"},
	}
	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			rest := MockRestricter{Data: map[string]string{"motion/1/field_ids": tt.value}}
			keychanges := newMockKeyChanged()
			defer keychanges.close()
			s := autoupdate.New(rest, keychanges)
			ider := s.IDer(1)

			_, err := ider.IDs(context.Background(), "motion/1/field_ids")

			if err == nil {
				t.Fatalf("Expected an error, got None")
			}
			if got := err.Error(); !strings.HasPrefix(got, tt.err) {
				t.Errorf("Expected error msg to be `%s`, got `%s`", tt.err, got)
			}
			return
		})
	}
}
