package dskey_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

func TestFromString(t *testing.T) {
	for _, tt := range []struct {
		key   string
		valid bool
	}{
		{"user/1/username", true},
		{"user/1/us.ername", false},
		{"user/1username", false},
		{"user/einz/ername", false},
	} {
		t.Run(tt.key, func(t *testing.T) {
			key, err := dskey.FromString(tt.key)
			if tt.valid && err != nil {
				t.Errorf("Key is not valid, expected valid key")
			} else if !tt.valid && err == nil {
				t.Errorf("Key is valid, expected invalid key")
			}

			if tt.valid && key.String() != tt.key {
				t.Errorf("build key != created key: %s != %s", tt.key, key.String())
			}
		})
	}
}

func TestFromParts(t *testing.T) {
	for _, tt := range []struct {
		collection string
		id         int
		field      string
		fromString string
	}{
		{"user", 1, "username", "user/1/username"},
		{"user", 12, "username", "user/12/username"},
		{"motion_version", 12, "username", "motion_version/12/username"},
	} {
		t.Run(fmt.Sprintf("%s/%d/%s", tt.collection, tt.id, tt.field), func(t *testing.T) {
			keyFromParts := dskey.FromParts(tt.collection, tt.id, tt.field)
			keyFromString, _ := dskey.FromString(tt.fromString)

			if keyFromParts != keyFromString {
				t.Errorf("from parts != from string")
			}
		})
	}
}

func TestID(t *testing.T) {
	for _, tt := range []struct {
		key    string
		expect int
	}{
		{"user/1/username", 1},
		{"user/12/username", 12},
	} {
		t.Run(tt.key, func(t *testing.T) {
			key, err := dskey.FromString(tt.key)
			if err != nil {
				t.Fatalf("Key is not valid: %v", err)
			}

			if key.ID() != tt.expect {
				t.Errorf("got %d, expected %d", key.ID(), tt.expect)
			}
		})
	}
}

func TestCollection(t *testing.T) {
	for _, tt := range []struct {
		key    string
		expect string
	}{
		{"user/1/username", "user"},
		{"user/12/username", "user"},
		{"motion_version/12/username", "motion_version"},
	} {
		t.Run(tt.key, func(t *testing.T) {
			key, err := dskey.FromString(tt.key)
			if err != nil {
				t.Fatalf("Key is not valid: %v", err)
			}

			if key.Collection() != tt.expect {
				t.Errorf("got %s, expected %s", key.Collection(), tt.expect)
			}
		})
	}
}

func TestField(t *testing.T) {
	for _, tt := range []struct {
		key    string
		expect string
	}{
		{"user/1/username", "username"},
		{"user/12/username", "username"},
		{"motion_version/12/with_$_template", "with_$_template"},
	} {
		t.Run(tt.key, func(t *testing.T) {
			key, err := dskey.FromString(tt.key)
			if err != nil {
				t.Fatalf("Key is not valid: %v", err)
			}

			if key.Field() != tt.expect {
				t.Errorf("got %s, expected %s", key.Field(), tt.expect)
			}
		})
	}
}

func TestFQID(t *testing.T) {
	for _, tt := range []struct {
		key    string
		expect string
	}{
		{"user/1/username", "user/1"},
		{"user/12/username", "user/12"},
		{"motion_version/12/with_$_template", "motion_version/12"},
	} {
		t.Run(tt.key, func(t *testing.T) {
			key, err := dskey.FromString(tt.key)
			if err != nil {
				t.Fatalf("Key is not valid: %v", err)
			}

			if key.FQID() != tt.expect {
				t.Errorf("got %s, expected %s", key.FQID(), tt.expect)
			}
		})
	}
}

func TestCollectionField(t *testing.T) {
	for _, tt := range []struct {
		key    string
		expect string
	}{
		{"user/1/username", "user/username"},
		{"user/12/username", "user/username"},
		{"motion_version/12/with_$_template", "motion_version/with_$_template"},
	} {
		t.Run(tt.key, func(t *testing.T) {
			key, err := dskey.FromString(tt.key)
			if err != nil {
				t.Fatalf("Key is not valid: %v", err)
			}

			if key.CollectionField() != tt.expect {
				t.Errorf("got %s, expected %s", key.CollectionField(), tt.expect)
			}
		})
	}
}

func TestIDField(t *testing.T) {
	for _, tt := range []struct {
		key    string
		expect string
	}{
		{"user/1/username", "user/1/id"},
		{"user/12/username", "user/12/id"},
		{"motion_version/12/with_$_template", "motion_version/12/id"},
		{"motion_version/12/with_$2_template", "motion_version/12/id"},
	} {
		t.Run(tt.key, func(t *testing.T) {
			key, err := dskey.FromString(tt.key)
			if err != nil {
				t.Fatalf("Key is not valid: %v", err)
			}

			if key.IDField().String() != tt.expect {
				t.Errorf("got %s, expected %s", key.IDField(), tt.expect)
			}
		})
	}
}

func TestReplaceField(t *testing.T) {
	key := dskey.MustKey("user/1/group_$_ids")
	value := "2"
	newkey, _ := dskey.FromString(key.FQID() + "/" + strings.Replace(key.Field(), "$", "$"+value, 1))

	expect := "user/1/group_$2_ids"
	if got := newkey.String(); got != expect {
		t.Errorf("got %s, expected %s", got, expect)
	}
}
