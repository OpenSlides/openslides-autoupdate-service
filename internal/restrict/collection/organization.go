package collection

// Organization handels restrictions of the collection organization.
//
// The user can always see an organization.
//
// Mode A: The user can see the organization (always).
//
// Mode B: The user must be logged in (no anonymous).
type Organization struct{}

// Modes returns the restrictions modes for the meeting collection.
func (o Organization) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return Allways
	case "B":
		return loggedIn
	}
	return nil
}
