package collection

// OrganizationTag handels restrictions of the collection organization_tag.
type OrganizationTag struct{}

// Modes returns the restrictions modes for the meeting collection.
func (o OrganizationTag) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return loggedIn
	}
	return nil
}
