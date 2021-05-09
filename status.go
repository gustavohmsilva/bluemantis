package bluemantis

// Status holds a given status from MantisBT. Default Status are normally:
// "new", "feedback", "acknowledged", "confirmed", "assigned", "resolved" and
// "closed".
type Status struct {
	Rel
	Color string `valid:"-" json:"color"`
}
