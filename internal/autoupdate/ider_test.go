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
	s := autoupdate.New(test.MockRestricter{}, keychanges)
	defer s.Close()
	ider := s.IDer(1)

	tc := []struct {
		name    string
		key     string
		idCount int
		err     string
	}{
		{"Restricter error", "error_id", 0, "can not restrict key error_id:"},
		{"IDs field", "motion/1/category_ids", 2, ""},
	}
	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			ids, err := ider.IDList(context.Background(), tt.key)
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
			rest := test.MockRestricter{Data: keyValue{"motion/1/field_ids": tt.value}.m()}
			keychanges := test.NewMockKeysChanged()
			defer keychanges.Close()
			s := autoupdate.New(rest, keychanges)
			defer s.Close()
			ider := s.IDer(1)

			_, err := ider.IDList(context.Background(), "motion/1/field_ids")

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
