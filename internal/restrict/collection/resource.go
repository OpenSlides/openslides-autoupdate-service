package collection

// Resource handels restrictions of the collection resource.
//
// The user can always see a resource.
//
// Mode A: The user can see the resource (always).
type Resource struct{}

// Modes returns the restrictions modes for the meeting collection.
func (r Resource) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return Allways
	}
	return nil
}
