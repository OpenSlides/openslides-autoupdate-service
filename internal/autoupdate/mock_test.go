package autoupdate_test

type mockKeysBuilder struct {
	keys []string
}

func (m mockKeysBuilder) Update([]string) error {
	return nil
}

func (m mockKeysBuilder) Keys() []string {
	return m.keys
}

// keyValue is a helper to create a map from string to string.
type keyValue map[string]string

func (kv keyValue) m() map[string]string {
	out := make(map[string]string)
	for key, value := range kv {
		out[key] = value
	}
	return out
}

// keys is a helper function to create a slice of strings.
func keys(keys ...string) []string { return keys }
