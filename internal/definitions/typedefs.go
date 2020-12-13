package definitions

import "encoding/json"

// ID TODO
type ID = int

// Fqid TODO
type Fqid = string

// Collection TODO
type Collection = string

// Fqfield TODO
type Fqfield = string

// Value TODO
type Value = json.RawMessage

// Field TODO
type Field = string

// FqfieldData TODO
type FqfieldData = map[Fqfield]Value

// Addition TODO
type Addition = map[string]interface{}
