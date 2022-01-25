package collection

// Theme handels the restrictions for the theme collection.
//
// Every user can see a theme.
type Theme struct{}

// Modes returns the field restriction for each mode.
func (t Theme) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return Allways
	}
	return nil
}
