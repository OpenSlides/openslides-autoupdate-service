package dskey

// Relation tells if a key relates somewhere else.
type Relation int

// This a the different types of relations.
const (
	RelationNone Relation = iota
	RelationSingle
	RelationList
	RelationGenericSingle
	RelationGenericList
)
