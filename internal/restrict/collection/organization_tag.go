package collection

// OrganizationTag handels restrictions of the collection organization_tag.
//
// A logged in user can always see an organization tag.
//
// Mode A: The user can see the organization tag.
type OrganizationTag struct{}

// Modes returns the restrictions modes for the meeting collection.
func (o OrganizationTag) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return loggedIn
	}
	return nil
}
