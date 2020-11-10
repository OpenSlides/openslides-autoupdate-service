package definitions

import "encoding/json"

type Id = int
type Fqid = string
type Fqfield = string
type Value = json.RawMessage
type Field = string
type FqfieldData = map[Fqfield]Value
type Addition = map[string]interface{}
