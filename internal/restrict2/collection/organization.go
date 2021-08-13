package collection

// Organization handels restrictions of the collection organization.
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
