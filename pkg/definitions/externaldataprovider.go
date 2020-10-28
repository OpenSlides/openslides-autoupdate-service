package definitions

type ExternalDataProvider interface {
	// If a field does not exist, it is not returned.
	Get(fqfields []Fqfield) FqfieldData
}
