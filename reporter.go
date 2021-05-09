package bluemantis

// Reporter represents who is reporting a given error. It usually matches the
// user who generated the Authorization Token of the application, but may be
// different depending on the configuration on MantisBT.
type Reporter struct {
	Rel
	Email string `valid:"email,optional" json:"email"`
}
