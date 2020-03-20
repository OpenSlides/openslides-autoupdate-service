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

// keys is a helper function to create a slice of strings.
func keys(keys ...string) []string { return keys }
