package dskey

// MetaKey can contain a key for meta information.
type MetaKey struct {
	key Key
	str string
}

// MetaFromKey returns a meta key from a key.
func MetaFromKey(key Key) MetaKey {
	return MetaKey{
		key: key,
	}
}

// MetaFromStr returns a Meta from a string.
func MetaFromStr(str string) MetaKey {
	return MetaKey{
		str: str,
	}
}

// Key returns the key, if the MetaKey was created with MetaFromKey
func (m MetaKey) Key() (Key, bool) {
	if m.key != 0 {
		return m.key, true
	}
	return 0, false
}

// Str returns the str, if the MetaKey was created with MetaFromStr
func (m MetaKey) Str() (string, bool) {
	if m.key == 0 {
		return m.str, true
	}
	return "", false
}

// KeyStr returns the Key and the str. The third value is true, if its a key.
func (m MetaKey) KeyStr() (Key, string, bool) {
	return m.key, m.str, m.key != 0
}
